import functools
import logging
import os

import grpc
import pika

from models import Message
from verificator import ModelVerificator
from ardea_pb2_grpc import ArdeaStub

LOG_FORMAT = '%(levelname)s\t%(asctime)s\t%(name)s\t:\t%(message)s'

logger = logging.getLogger(__name__)

logging.basicConfig(level=logging.INFO, format=LOG_FORMAT)


def ack_message(channel, delivery_tag):
    if channel.is_open:
        channel.basic_ack(delivery_tag)


def on_message(channel, method_frame, header_frame, body, args):
    connection, verificator = args

    delivery_tag = method_frame.delivery_tag

    logger.info(f'processing message {delivery_tag}: {body}')
    try:
        msg = Message.parse_raw(body)
        verificator.verify(msg)
    except Exception as e:
        logger.error(f"error handling model verification: {repr(e)}")

    logger.info(f'message {delivery_tag} processed')
    cb = functools.partial(ack_message, channel, delivery_tag)
    connection.add_callback_threadsafe(cb)


if __name__ == '__main__':
    grpc_channel = grpc.insecure_channel(os.getenv('ARDEA_TARGET'))
    ardea_stub = ArdeaStub(grpc_channel)

    verificator = ModelVerificator(ardea_stub, os.getenv('MODELS_STORE_TARGET'))

    connection = pika.BlockingConnection(pika.ConnectionParameters(os.getenv('RABBITMQ_TARGET'), heartbeat=5))

    channel = connection.channel()
    channel.queue_declare(queue="models_verify")

    on_message_callback = functools.partial(on_message, args=(connection, verificator))
    channel.basic_consume(queue='models_verify', on_message_callback=on_message_callback, auto_ack=False)

    try:
        channel.start_consuming()
    except KeyboardInterrupt:
        channel.stop_consuming()

    connection.close()
