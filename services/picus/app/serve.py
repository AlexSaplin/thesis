import grpc_source.picus_pb2 as pb
import grpc_source.picus_pb2_grpc
from google.protobuf.timestamp_pb2 import Timestamp

import asyncio

from kafka_reader import read_topic


class PicusServicer(grpc_source.picus_pb2_grpc.PicusServicer):

    def __init__(self, kafka_broker, logger):
        self.logger = logger
        self.kafka_broker = kafka_broker

    async def GetFunctionLogs(self, request, context):
        logs = []
        try:
            async for item in read_topic(self.kafka_broker, request.function_id):
                timestamp = Timestamp()
                timestamp.seconds = item["timestamp"]
                logs.append(pb.LogEntry(
                    function_id=request.function_id,
                    time=timestamp,
                    message=item["log"]
                ))

            self.logger.info(f'Successfully fetched logs for function_id: {request.function_id}')
        except asyncio.CancelledError:
            self.logger.info(f'Cancelled fetching function logs for function_id: {request.function_id}')
            return pb.GetFunctionLogsReply(entries=logs)
        except Exception as e:
            self.logger.error(
                f'Failed to fetch function logs for function_id: {request.function_id}. Error: {repr(e)}')
            raise
        return pb.GetFunctionLogsReply(entries=logs)

    async def StreamFunctionLogs(self, request, context):
        try:
            async for item in read_topic(self.kafka_broker, request.function_id, stream=True):
                timestamp = Timestamp()
                timestamp.seconds = item["timestamp"]
                self.logger.info(f'Successfully fetched log entry for function_id: {request.function_id}')
                yield pb.LogEntry(
                    function_id=request.function_id,
                    time=timestamp,
                    message=item["log"]
                )
        except asyncio.CancelledError:
            self.logger.info(f'Cancelled fetching function logs for function_id: {request.function_id}')
            return
        except Exception as e:
            self.logger.error(f'Failed to fetch function logs for function_id: {request.function_id}. Error: {repr(e)}')
            raise
