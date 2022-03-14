import base64

import grpc

import grpc_source.slav_pb2_grpc
import grpc_source.slav_pb2 as pb
from db.adapter import SlavAdapter
from db.interface import DatabaseInterface
from errors import ContainerAlreadyExistsException
from tesseract.client import TesseractClient


class SlavServicer(grpc_source.slav_pb2_grpc.SlavServicer):

    def __init__(self, db_url: str, stub: TesseractClient, logger):
        self.db = DatabaseInterface(db_url=db_url, tesseract_client=stub, logger=logger)
        self.logger = logger

    def CreateContainer(self, request, context):
        try:
            env = [{"key": item.Key, "value": item.Value} for item in request.Env]

            auth = {
                "auths": {
                    item.Registry: {
                        "username": item.Username,
                        "password": item.Password,
                        "auth": base64.b64encode(f'{item.Username}:{item.Password}'.encode()).decode()
                    }
                    for item in request.Auth
                }
            }

            container = self.db.create_container(
                owner_id=request.OwnerID,
                name=request.Name,
                scale=request.Scale,
                instance_type=SlavAdapter.deserialize_instance_type(request.Instance),
                image=request.Image,
                port=request.Port,
                env=env,
                auth=auth,
            )
            self.logger.info(f'Container creation successfully started. '
                             f'Name: {request.Name}, '
                             f'OwnerID: {request.OwnerID}, '
                             f'Image: {request.Image}')
            container_proto = SlavAdapter.container_dataclass_to_protobuf(container)
        except ContainerAlreadyExistsException:
            context.set_code(grpc.StatusCode.ALREADY_EXISTS)
            context.set_details(f'Container with name {request.Name} already exist')
            return pb.CreateContainerReply()
        except Exception as e:
            self.logger.error(f'Container creation failed '
                              f'Name: {request.Name}, '
                              f'OwnerID: {request.OwnerID}, '
                              f'Image: {request.Image}. Error: {repr(e)}')
            raise

        return pb.CreateContainerReply(
            Container=container_proto,
        )

    def UpdateContainer(self, request, context):
        try:
            params = {}

            if request.Scale.IsValid:
                params['scale'] = request.Scale.Value

            if request.Instance.IsValid:
                params['instance_type'] = SlavAdapter.deserialize_instance_type(request.Instance.Value).value

            if request.Image.IsValid:
                params['image'] = request.Image.Value

            container = self.db.update_container(
                owner_id=request.OwnerID,
                name=request.Name,
                params=params
            )
            container = SlavAdapter.container_dataclass_to_protobuf(container)
            self.logger.info(f'Container {request.Name} is updating...')
        except Exception as e:
            self.logger.error(f'Failed to update container {request.Name}. Error: {repr(e)}')
            raise

        return pb.CreateContainerReply(
            Container=container,
        )

    def GetContainer(self, request, context):
        try:
            container, status, error = self.db.get_container(
                owner_id=request.OwnerID,
                name=request.Name,
            )
            self.logger.info(f'Container fetched successfully {request.Name}')
            container = SlavAdapter.container_dataclass_to_protobuf(container)
        except Exception as e:
            self.logger.error(f'Failed to get container {request.Name}. Error: {repr(e)}')
            raise

        return pb.GetContainerReply(
            Container=container,
            State=SlavAdapter.serialize_state_type(status),
            Error=error,
        )

    def DeleteContainer(self, request, context):
        try:
            self.db.delete_container(
                owner_id=request.OwnerID,
                name=request.Name,
            )
            self.logger.info(f'Container removed successfully {request.Name}')
        except Exception as e:
            self.logger.error(f'Failed to remove container {request.Name}. Error: {repr(e)}')
            raise

        return pb.DeleteContainerReply()

    def ListContainers(self, request, context):
        try:
            containers = self.db.list_containers(
                owner_id=request.OwnerID
            )
            containers = [SlavAdapter.container_dataclass_to_protobuf(container) for container in containers]
            self.logger.info(f'Containers, owned by OwnerID: {request.OwnerID}, fetched successfully')
        except Exception as e:
            self.logger.error(f'Failed to fetch containers, owned by OwnerID: {request.OwnerID}. Error: {repr(e)}')
            raise

        return pb.ListContainersReply(
            Containers=containers
        )
