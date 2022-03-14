import typing

from datetime import datetime

from gorilla_pb2 import AddDeltasRequest, AddDeltasResponse, Delta
from gorilla_pb2_grpc import GorillaStub

from gorilla_pb2 import GetDeltasResponse, GetDeltasRequest


def calc_price(current_usage: float):
    return 0.36 / 3600


def send_bill_to_gorilla(gorilla: GorillaStub, deltas: typing.List[Delta]):
    _: AddDeltasResponse = gorilla.AddDeltas(AddDeltasRequest(Deltas=deltas))


def make_delta(*, owner_id: str, object_id: str,
               balance: float, category: str, date: datetime) -> Delta:
    print(f'got delta owner_id={owner_id} object_id={object_id} balance={balance} category={category} date={date}')
    return Delta(
        Date=int(date.timestamp()),
        Category=category,
        Balance=balance,
        OwnerID=owner_id,
        ObjectID=object_id,
        ObjectType='FUNCTION',
    )
