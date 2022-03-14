import config
from db.models import ContainerMeta
import grpc_source.slav_pb2 as pb
import grpc_source.tesseract_pb2 as pb_t
from slav_dataclasses import Container, InstanceType, StateType, InstanceSpecType


class SlavAdapter:

    def __init__(self):
        ...

    @staticmethod
    def serialize_instance_type(instance_type: InstanceType) -> pb.InstanceType:
        if instance_type == InstanceType.STARTER:
            return pb.InstanceType.STARTER
        elif instance_type == InstanceType.INFERENCE:
            return pb.InstanceType.INFERENCE
        else:
            raise ValueError(f'Wrong instance ENUM value. Got: {getattr(instance_type, "value", None)}')

    @staticmethod
    def deserialize_instance_type(instance_type: pb.InstanceType) -> InstanceType:
        if instance_type == pb.InstanceType.STARTER:
            return InstanceType.STARTER
        elif instance_type == pb.InstanceType.INFERENCE:
            return InstanceType.INFERENCE
        else:
            raise ValueError(f'Wrong instance ENUM value. Got: {getattr(instance_type, "value", None)}')

    @staticmethod
    def serialize_state_type(state_type: StateType) -> pb.StateType:
        if state_type == StateType.RUNNING:
            return pb.StateType.RUNNING
        elif state_type == StateType.UPDATING:
            return pb.StateType.UPDATING
        elif state_type == StateType.ERROR:
            return pb.StateType.ERROR
        else:
            return pb.StateType.UNKNOWN

    @staticmethod
    def deserialize_state_type(state_type: pb.StateType) -> StateType:
        if state_type == pb.StateType.RUNNING:
            return StateType.RUNNING
        elif state_type == pb.StateType.UPDATING:
            return StateType.UPDATING
        elif state_type == pb.StateType.ERROR:
            return StateType.ERROR
        else:
            return StateType.UNKNOWN

    @staticmethod
    def container_meta_to_dataclass(container: ContainerMeta) -> Container:
        return Container(
            ID=str(container.id),
            Name=container.name,
            Scale=container.scale,
            Instance=InstanceType(container.instance_type),
            Image=container.image,
            Port=container.port,
            OwnerID=container.owner_id,
            URL=f'http://{container.name}.{config.CONTAINERS_DOMAIN}',
            Env=container.env,
            Auth=container.auth
        )

    @staticmethod
    def container_dataclass_to_protobuf(container: Container) -> pb.Container:
        env = []
        if container.Env:
            env = [pb.KV(Key=item["key"], Value=item["value"]) for item in container.Env]
        auth = []
        if container.Auth and "auths" in container.Auth:
            auth = list(container.Auth["auths"].keys())
        return pb.Container(
            Name=container.Name,
            Scale=container.Scale,
            Instance=SlavAdapter.serialize_instance_type(container.Instance),
            Image=container.Image,
            Port=container.Port,
            URL=container.URL,
            Env=env,
            Auth=auth,
        )

    @staticmethod
    def tesseract_status_to_dataclass(status: pb_t.Status) -> StateType:
        if status == pb_t.RUNNING:
            return StateType.RUNNING
        elif status == pb_t.UPDATING:
            return StateType.UPDATING
        elif status == pb_t.ERROR:
            return StateType.ERROR
        else:
            return StateType.UNKNOWN

    @staticmethod
    def instance_type_to_spec_class(instance_type: InstanceType) -> InstanceSpecType:
        if instance_type == InstanceType.STARTER:
            return InstanceSpecType.STARTER
        elif instance_type == InstanceType.INFERENCE:
            return InstanceSpecType.INFERENCE
        else:
            raise ValueError(f'Wrong instance ENUM value. Got: {getattr(instance_type, "value", None)}')

    @staticmethod
    def instance_type_to_billing_price(instance_type: InstanceType) -> config.BillingPrice:
        if instance_type == InstanceType.STARTER:
            return config.BillingPrice.STARTER_PRICE
        elif instance_type == InstanceType.INFERENCE:
            return config.BillingPrice.INFERENCE_PRICE
        else:
            raise ValueError(f'Wrong instance ENUM value. Got: {getattr(instance_type, "value", None)}')