# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import clients.ibis.ibis_pb2 as ibis__pb2


class IbisStub(object):
    """The Ibis service definition.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateFunction = channel.unary_unary(
                '/ibis.Ibis/CreateFunction',
                request_serializer=ibis__pb2.CreateFunctionRequest.SerializeToString,
                response_deserializer=ibis__pb2.CreateFunctionReply.FromString,
                )
        self.GetFunction = channel.unary_unary(
                '/ibis.Ibis/GetFunction',
                request_serializer=ibis__pb2.GetFunctionRequest.SerializeToString,
                response_deserializer=ibis__pb2.GetFunctionReply.FromString,
                )
        self.GetFunctionByName = channel.unary_unary(
                '/ibis.Ibis/GetFunctionByName',
                request_serializer=ibis__pb2.GetFunctionByNameRequest.SerializeToString,
                response_deserializer=ibis__pb2.GetFunctionReply.FromString,
                )
        self.UpdateFunction = channel.unary_unary(
                '/ibis.Ibis/UpdateFunction',
                request_serializer=ibis__pb2.UpdateFunctionRequest.SerializeToString,
                response_deserializer=ibis__pb2.UpdateFunctionReply.FromString,
                )
        self.ListFunctions = channel.unary_unary(
                '/ibis.Ibis/ListFunctions',
                request_serializer=ibis__pb2.ListFunctionsRequest.SerializeToString,
                response_deserializer=ibis__pb2.ListFunctionsReply.FromString,
                )


class IbisServicer(object):
    """The Ibis service definition.
    """

    def CreateFunction(self, request, context):
        """Missing associated documentation comment in .proto file"""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetFunction(self, request, context):
        """Missing associated documentation comment in .proto file"""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetFunctionByName(self, request, context):
        """Missing associated documentation comment in .proto file"""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateFunction(self, request, context):
        """Missing associated documentation comment in .proto file"""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListFunctions(self, request, context):
        """Missing associated documentation comment in .proto file"""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_IbisServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateFunction': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateFunction,
                    request_deserializer=ibis__pb2.CreateFunctionRequest.FromString,
                    response_serializer=ibis__pb2.CreateFunctionReply.SerializeToString,
            ),
            'GetFunction': grpc.unary_unary_rpc_method_handler(
                    servicer.GetFunction,
                    request_deserializer=ibis__pb2.GetFunctionRequest.FromString,
                    response_serializer=ibis__pb2.GetFunctionReply.SerializeToString,
            ),
            'GetFunctionByName': grpc.unary_unary_rpc_method_handler(
                    servicer.GetFunctionByName,
                    request_deserializer=ibis__pb2.GetFunctionByNameRequest.FromString,
                    response_serializer=ibis__pb2.GetFunctionReply.SerializeToString,
            ),
            'UpdateFunction': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateFunction,
                    request_deserializer=ibis__pb2.UpdateFunctionRequest.FromString,
                    response_serializer=ibis__pb2.UpdateFunctionReply.SerializeToString,
            ),
            'ListFunctions': grpc.unary_unary_rpc_method_handler(
                    servicer.ListFunctions,
                    request_deserializer=ibis__pb2.ListFunctionsRequest.FromString,
                    response_serializer=ibis__pb2.ListFunctionsReply.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'ibis.Ibis', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Ibis(object):
    """The Ibis service definition.
    """

    @staticmethod
    def CreateFunction(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ibis.Ibis/CreateFunction',
            ibis__pb2.CreateFunctionRequest.SerializeToString,
            ibis__pb2.CreateFunctionReply.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetFunction(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ibis.Ibis/GetFunction',
            ibis__pb2.GetFunctionRequest.SerializeToString,
            ibis__pb2.GetFunctionReply.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetFunctionByName(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ibis.Ibis/GetFunctionByName',
            ibis__pb2.GetFunctionByNameRequest.SerializeToString,
            ibis__pb2.GetFunctionReply.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateFunction(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ibis.Ibis/UpdateFunction',
            ibis__pb2.UpdateFunctionRequest.SerializeToString,
            ibis__pb2.UpdateFunctionReply.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListFunctions(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ibis.Ibis/ListFunctions',
            ibis__pb2.ListFunctionsRequest.SerializeToString,
            ibis__pb2.ListFunctionsReply.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)