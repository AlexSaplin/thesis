from concurrent import futures
import logging
import asyncio

from grpc_source.picus_pb2_grpc import add_PicusServicer_to_server
from serve import PicusServicer
import grpc
import config

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)


async def serve():
    server = grpc.aio.server(futures.ThreadPoolExecutor(max_workers=5))

    port = config.PORT
    kafka_broker = config.KAFKA_BROKER

    add_PicusServicer_to_server(PicusServicer(kafka_broker, logger), server)

    server.add_insecure_port('[::]:{}'.format(port))
    await server.start()
    await server.wait_for_termination()
    logger.info(f'Started server on [::]:{port}')


if __name__ == '__main__':
    asyncio.run(serve())
