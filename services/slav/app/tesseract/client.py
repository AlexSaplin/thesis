import base64
import json

from db.adapter import SlavAdapter
from grpc_source.tesseract_pb2_grpc import TesseractStub
import grpc_source.tesseract_pb2 as pb
from slav_dataclasses import InstanceType, StateType


class TesseractClient:

    def __init__(self, stub: TesseractStub, containers_domain: str, logger):
        self.stub = stub
        self.containers_domain = containers_domain
        self.logger = logger

    def apply_changes(self, name: str, owner_id: str, scale: int, port: int, image: str, instance_type: InstanceType, env: list, auth: dict):
        try:
            instance_spec = SlavAdapter.instance_type_to_spec_class(instance_type)
            env_pb = [pb.KV(Key=item["key"], Value=item["value"]) for item in env]
            auth_pb = base64.b64encode(json.dumps(auth).encode()).decode()
            req = pb.ApplyRequest(
                Name=name,
                ID=f'{name}-{owner_id}',
                DNS=f'{name}.{self.containers_domain}',
                Scale=scale,
                CPU=instance_spec.value.cpu,
                RAM=instance_spec.value.ram,
                GPU=instance_spec.value.gpu,
                Port=port,
                Image=image,
                Env=env_pb,
                Auth=auth_pb,
            )
            self.stub.Apply(req)
            self.logger.info(f'Changes have been applied successfully. Name: {name}, OwnerID: {owner_id}')
        except Exception as e:
            self.logger.error(f'Failed to apply changes on container. Name: {name}, OwnerID: {owner_id}. '
                              f'Error: {repr(e)}')
            raise

    def delete_container(self, name: str, owner_id: str):
        try:
            req = pb.DeleteRequest(
                ID=f'{name}-{owner_id}'
            )
            self.stub.Delete(req)
            self.logger.info(f'Container removed successfully {name}-{owner_id}')
        except Exception as e:
            self.logger.error(f'Failed to remove container {name}-{owner_id}. Error: {repr(e)}')
            raise

    def get_status(self, name: str, owner_id: str):
        error = None
        try:
            req = pb.GetStatusRequest(
                ID=f'{name}-{owner_id}'
            )
            resp = self.stub.GetStatus(req)
            result = SlavAdapter.tesseract_status_to_dataclass(resp.Status)
            if resp.Error:
                result = StateType.ERROR
                error = resp.Error
            self.logger.info(f'Container status fetched successfully {name}-{owner_id}')
        except Exception as e:
            self.logger.error(f'Failed to fetch status of container {name}-{owner_id}. Error: {repr(e)}')
            raise
        return result, error
