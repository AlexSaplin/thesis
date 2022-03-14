"""
Async kafka reader that handles consumer management.
"""
import asyncio

import aiokafka
from kafka import TopicPartition

import json


def serialize(x):
    if x is None:
        return x
    return json.dumps(x).encode('utf-8')


def deserialize(x):
    if x is None:
        return x
    return json.loads(x.decode('utf-8'))


try:  # Python 3.7
    from contextlib import asynccontextmanager
except ImportError:  # Python 3.6
    from async_generator import asynccontextmanager


@asynccontextmanager
async def kafka_consumer(bootstrap_servers, topic_name):
    loop = asyncio.get_event_loop()
    consumer = aiokafka.AIOKafkaConsumer(
        topic_name,
        bootstrap_servers=bootstrap_servers,
        key_deserializer=deserialize,
        value_deserializer=deserialize,
        loop=loop,
    )
    try:
        await consumer.start()
        yield consumer
    finally:
        await consumer.stop()


async def read_topic(bootstrap_servers, topic_name, start=None, stop=None, stream=False):
    """
    Read topic starting from `start` up until `stop`. If `stop` is not provided, continue reading forever.
    :param topic_name: Topic to read from
    :param start:
    :param stop:
    :param amount:
    :param stream:
    :return:
    """
    async with kafka_consumer(bootstrap_servers, topic_name) as consumer:
        partitions = consumer.partitions_for_topic(topic_name)
        if not partitions:
            raise RuntimeError(f"No partitions for topic {topic_name}")
        partition = next(iter(partitions))
        partition = TopicPartition(topic=topic_name, partition=partition)
        # Fetch some data to get highwater from kafka
        consumer.seek(partition, 0)
        try:
            await asyncio.wait_for(consumer.getone(partition), 30.0)
        except asyncio.TimeoutError:
            return
        size = consumer.highwater(partition)

        start = start if start is not None else 0
        start = max(0, start if start >= 0 else size + start)

        stop = stop if stop is not None else size
        stop = max(0, stop if stop >= 0 else size + stop)
        if not stream and stop <= start:
            return
        consumer.seek(partition, start)
        last_offset = -1
        while stream or last_offset < stop - 1:
            record = await consumer.getone(partition)
            last_offset = record.offset
            yield record.value
