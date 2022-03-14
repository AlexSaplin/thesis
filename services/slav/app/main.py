import logging
import signal
import time
from concurrent import futures

import grpc
from grpc_reflection.v1alpha import reflection
import schedule

import config
from billing_job import job
from db.interface import DatabaseInterface
from gorilla.client import GorillaClient
from grpc_source.gorilla_pb2_grpc import GorillaStub
from grpc_source.tesseract_pb2_grpc import TesseractStub
from serve import SlavServicer
from grpc_source.slav_pb2 import DESCRIPTOR
from grpc_source.slav_pb2_grpc import add_SlavServicer_to_server
from tesseract.client import TesseractClient

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)


class GracefulKiller:
    __kill_now = False

    def __init__(self):
        signal.signal(signal.SIGINT, self.exit_gracefully)
        signal.signal(signal.SIGTERM, self.exit_gracefully)

    @classmethod
    def exit_gracefully(cls, signum, frame):
        logger.warning(f'Catched stop signal, gracefully exiting...', signum=signum, frame=frame)
        cls.__kill_now = True

    @classmethod
    def not_now(cls):
        return not cls.__kill_now


def serve():
    logger.info(f'Starting Slav server')

    killer = GracefulKiller()

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=3))

    db_url = config.DB_URL
    port = config.SLAV_PORT
    containers_domain = config.CONTAINERS_DOMAIN

    grpc_channel = grpc.insecure_channel(config.TESSERACT_TARGET)
    tesseract_stub = TesseractStub(grpc_channel)
    tesseract_client = TesseractClient(tesseract_stub, containers_domain, logger)

    grpc_gorilla_channel = grpc.insecure_channel(config.GORILLA_TARGET)
    gorilla_stub = GorillaStub(grpc_gorilla_channel)
    gorilla_client = GorillaClient(gorilla_stub, logger)

    add_SlavServicer_to_server(SlavServicer(db_url, tesseract_client, logger), server)

    service_names = (
        DESCRIPTOR.services_by_name['Slav'].full_name,
        reflection.SERVICE_NAME,
    )
    reflection.enable_server_reflection(service_names, server)

    server.add_insecure_port(f'[::]:{port}')
    server.start()

    logger.info(f'Started server on [::]:{port}')

    db_interface = DatabaseInterface(db_url=db_url, tesseract_client=tesseract_client, logger=logger)
    schedule.every().minute.do(job,
                               db_interface=db_interface,
                               tesseract_client=tesseract_client,
                               gorilla_client=gorilla_client,
                               logger=logger)

    try:
        while killer.not_now():
            schedule.run_pending()
            time.sleep(1)
        server.stop(0)
    except KeyboardInterrupt:
        server.stop(0)
    finally:
        logger.info('Stopped Slav server')


if __name__ == '__main__':
    serve()
