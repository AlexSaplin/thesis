import os
import typing
from uuid import uuid1
from datetime import datetime, timedelta
from gorilla_pb2 import AddDeltasRequest, Delta
from utils import calc_price


def gen_add_deltas_for_owner(owner_id: str, n_models: int = 1,
                             dt_begin=datetime.utcnow() - timedelta(days=1),
                             dt_end=datetime.utcnow(),
                             inference_time: float = 0.025, inference_dev: float = 0.5,
                             density: float = 0.001, density_dev: float = 0.5) -> typing.List[AddDeltasRequest]:
    # dynamic import random
    random: typing.Any = __import__('random')
    random.seed(os.urandom(32))

    model_ids = list()
    for _ in range(n_models):
        model_ids.append(str(uuid1()))

    deltas = list()

    time_range: timedelta = (dt_end - dt_begin)

    estimate_model_runs_cnt: int = int(time_range * density / timedelta(seconds=inference_time))
    base_step: float = time_range * (1 - density) / timedelta(seconds=estimate_model_runs_cnt)

    current_usage = 0.

    for model_id in model_ids:
        dt_cur = dt_begin
        while dt_cur < dt_end:
            inference_time_cur = inference_time + random.uniform(0, inference_dev)
            balance = -calc_price(current_usage) * inference_time_cur
            print(model_id)
            deltas.append(Delta(
                Date=int(dt_cur.astimezone().timestamp()),
                Category='FUNCTIONS',
                Balance=balance,
                OwnerID=owner_id,
                ObjectID=model_id,
                ObjectType='FUNCTION',
            ))

            dt_cur += timedelta(seconds=(random.uniform(0, density_dev) + base_step))

    return AddDeltasRequest(Deltas=deltas)


if __name__ == '__main__':
    print(gen_add_deltas_for_owner('a9bb6d83-6d2f-4842-8d53-a37345d3fd3b'))
