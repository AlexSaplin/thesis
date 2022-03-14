# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import ardea_pb.ardea_pb2 as ardea__pb2


class ArdeaStub(object):
    """type ArdeaService interface {
    	CreateModel(ctx context.Context, OwnerID string, InputShape, OutputShape []uint64) (model entities.Model, err error)
    	GetModel(ctx context.Context, modelID string) (entities.Model, error)
    	UpdateModelState(ctx context.Context, modelID string, state entities.ModelState) (entities.Model, error)
    	UpdateModelPath(ctx context.Context, modelID, path string) (entities.Model, error)
    }

    The Ardea service definition.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateModel = channel.unary_unary(
                '/ardea.Ardea/CreateModel',
                request_serializer=ardea__pb2.CreateModelRequest.SerializeToString,
                response_deserializer=ardea__pb2.CreateModelReply.FromString,
                )
        self.GetModel = channel.unary_unary(
                '/ardea.Ardea/GetModel',
                request_serializer=ardea__pb2.GetModelRequest.SerializeToString,
                response_deserializer=ardea__pb2.GetModelReply.FromString,
                )
        self.GetModelByName = channel.unary_unary(
                '/ardea.Ardea/GetModelByName',
                request_serializer=ardea__pb2.GetModelByNameRequest.SerializeToString,
                response_deserializer=ardea__pb2.GetModelReply.FromString,
                )
        self.UpdateModelState = channel.unary_unary(
                '/ardea.Ardea/UpdateModelState',
                request_serializer=ardea__pb2.UpdateModelStateRequest.SerializeToString,
                response_deserializer=ardea__pb2.UpdateModelStateReply.FromString,
                )
        self.UpdateModelPath = channel.unary_unary(
                '/ardea.Ardea/UpdateModelPath',
                request_serializer=ardea__pb2.UpdateModelPathRequest.SerializeToString,
                response_deserializer=ardea__pb2.UpdateModelPathReply.FromString,
                )
        self.ListModels = channel.unary_unary(
                '/ardea.Ardea/ListModels',
                request_serializer=ardea__pb2.ListModelsRequest.SerializeToString,
                response_deserializer=ardea__pb2.ListModelsReply.FromString,
                )


class ArdeaServicer(object):
    """type ArdeaService interface {
    	CreateModel(ctx context.Context, OwnerID string, InputShape, OutputShape []uint64) (model entities.Model, err error)
    	GetModel(ctx context.Context, modelID string) (entities.Model, error)
    	UpdateModelState(ctx context.Context, modelID string, state entities.ModelState) (entities.Model, error)
    	UpdateModelPath(ctx context.Context, modelID, path string) (entities.Model, error)
    }

    The Ardea service definition.
    """

    def CreateModel(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetModel(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetModelByName(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateModelState(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateModelPath(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListModels(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ArdeaServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateModel': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateModel,
                    request_deserializer=ardea__pb2.CreateModelRequest.FromString,
                    response_serializer=ardea__pb2.CreateModelReply.SerializeToString,
            ),
            'GetModel': grpc.unary_unary_rpc_method_handler(
                    servicer.GetModel,
                    request_deserializer=ardea__pb2.GetModelRequest.FromString,
                    response_serializer=ardea__pb2.GetModelReply.SerializeToString,
            ),
            'GetModelByName': grpc.unary_unary_rpc_method_handler(
                    servicer.GetModelByName,
                    request_deserializer=ardea__pb2.GetModelByNameRequest.FromString,
                    response_serializer=ardea__pb2.GetModelReply.SerializeToString,
            ),
            'UpdateModelState': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateModelState,
                    request_deserializer=ardea__pb2.UpdateModelStateRequest.FromString,
                    response_serializer=ardea__pb2.UpdateModelStateReply.SerializeToString,
            ),
            'UpdateModelPath': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateModelPath,
                    request_deserializer=ardea__pb2.UpdateModelPathRequest.FromString,
                    response_serializer=ardea__pb2.UpdateModelPathReply.SerializeToString,
            ),
            'ListModels': grpc.unary_unary_rpc_method_handler(
                    servicer.ListModels,
                    request_deserializer=ardea__pb2.ListModelsRequest.FromString,
                    response_serializer=ardea__pb2.ListModelsReply.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'ardea.Ardea', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Ardea(object):
    """type ArdeaService interface {
    	CreateModel(ctx context.Context, OwnerID string, InputShape, OutputShape []uint64) (model entities.Model, err error)
    	GetModel(ctx context.Context, modelID string) (entities.Model, error)
    	UpdateModelState(ctx context.Context, modelID string, state entities.ModelState) (entities.Model, error)
    	UpdateModelPath(ctx context.Context, modelID, path string) (entities.Model, error)
    }

    The Ardea service definition.
    """

    @staticmethod
    def CreateModel(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ardea.Ardea/CreateModel',
            ardea__pb2.CreateModelRequest.SerializeToString,
            ardea__pb2.CreateModelReply.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetModel(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ardea.Ardea/GetModel',
            ardea__pb2.GetModelRequest.SerializeToString,
            ardea__pb2.GetModelReply.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetModelByName(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ardea.Ardea/GetModelByName',
            ardea__pb2.GetModelByNameRequest.SerializeToString,
            ardea__pb2.GetModelReply.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateModelState(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ardea.Ardea/UpdateModelState',
            ardea__pb2.UpdateModelStateRequest.SerializeToString,
            ardea__pb2.UpdateModelStateReply.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateModelPath(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ardea.Ardea/UpdateModelPath',
            ardea__pb2.UpdateModelPathRequest.SerializeToString,
            ardea__pb2.UpdateModelPathReply.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListModels(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ardea.Ardea/ListModels',
            ardea__pb2.ListModelsRequest.SerializeToString,
            ardea__pb2.ListModelsReply.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)
