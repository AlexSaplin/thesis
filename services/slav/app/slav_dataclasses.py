from dataclasses import dataclass, asdict
from enum import Enum


class BaseDataclass:
    def to_dict(self):
        result = {k: v for k, v in asdict(self).items() if v is not None}
        return result

    @classmethod
    def from_dict(cls, d):
        return cls(**d)


class InstanceType(Enum):
    STARTER = 'STARTER'
    INFERENCE = 'INFERENCE'


class InstanceSpecBase:
    def __init__(self, cpu: int, ram: int, gpu: str):
        self.cpu = cpu
        self.ram = ram
        self.gpu = gpu


class InstanceSpecType(Enum):
    STARTER = InstanceSpecBase(1, 2048, 'none')
    INFERENCE = InstanceSpecBase(4, 12288, 't4')


class StateType(Enum):
    RUNNING = 'RUNNING'
    UPDATING = 'UPDATING'
    ERROR = 'ERROR'
    UNKNOWN = 'UNKNOWN'


@dataclass
class Container(BaseDataclass):
    ID: str
    Name: str
    Scale: int
    Instance: InstanceType
    Image: str
    Port: int
    OwnerID: str
    URL: str
    Env: list
    Auth: dict
