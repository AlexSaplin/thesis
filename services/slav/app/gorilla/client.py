import grpc_source.gorilla_pb2 as pb
from grpc_source.gorilla_pb2_grpc import GorillaStub


class GorillaClient:

    def __init__(self, stub: GorillaStub, logger):
        self.stub = stub
        self.logger = logger

    def add_deltas(self, deltas):
        try:
            self.logger.info(f'Sending deltas to Gorilla')
            req = pb.AddDeltasRequest(
                Deltas=[
                    pb.Delta(
                        Date=delta['date'],
                        Category=delta['category'],
                        Balance=delta['balance'],
                        OwnerID=delta['owner_id'],
                        ObjectID=delta['object_id'],
                        ObjectType=delta['object_type'],
                    )
                    for delta in deltas
                ]
            )
            self.stub.AddDeltas(req)
        except Exception as e:
            self.logger.error(f'Failed to send deltas to Gorilla. Error: {repr(e)}')
