import typing
from datetime import datetime
from collections import defaultdict

import grpc
import faust

from config import Config
from utils import calc_price, send_bill_to_gorilla, make_delta

from gorilla_pb2_grpc import GorillaStub
from gorilla_pb2 import Delta

config = Config()

app = faust.App(
    'ploti-nalogi',
    broker=config.kafka_broker,
    store='memory://'
)

usages = app.Table('daily_usages', default=float, partitions=1)
bills = defaultdict(float)


class Event(faust.Record):
    owner_id: str
    object_id: str
    run_duration: float
    load_duration: float
    ts: int

    def make_key_without_date(self):
        return f'{self.owner_id}:{self.object_id}'

    def make_key(self):
        dt = datetime.utcfromtimestamp(self.ts).strftime("%Y-%m-%d")
        return f'{self.make_key_without_date()}:{dt}'

    @staticmethod
    def parse_key(key: str):
        t = key.split(':')
        return {
            'owner_id': t[0],
            'object_id': t[1]
        }


input_topic = app.topic(config.input_topic, value_type=Event, partitions=1)


@app.agent(input_topic)
async def process(messages):
    async for event in messages:
        current_usage = usages[event.make_key()]
        current_price = calc_price(current_usage)
        usages[event.make_key()] += round(event.run_duration, 2)
        bills[event.make_key_without_date()] += round(event.run_duration, 2) * current_price

grpc_channel = grpc.insecure_channel(config.gorilla_target)
gorilla_stub = GorillaStub(grpc_channel)


def flush_bills():
    global bills
    bills = defaultdict(float)


@app.timer(interval=config.billing_frequency_seconds)
async def send_bills():
    deltas: typing.List[Delta] = list()
    for key in bills.keys():
        if bills[key] != 0:
            deltas.append(make_delta(
                **Event.parse_key(key),
                balance=-bills[key],
                category='FUNCTIONS',
                date=datetime.utcnow()
            ))
            bills[key] = 0
    if len(deltas) > 0:
        send_bill_to_gorilla(gorilla=gorilla_stub, deltas=deltas)
        flush_bills()
