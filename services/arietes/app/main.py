import asyncio
import logging

import faust

from builder import FunctionBuilder
from config import KAFKA_BROKER, INPUT_TOPIC


class FunctionQuery(faust.Record):
    function_id: str


app = faust.App('function-build-events', broker=KAFKA_BROKER, store='memory://')

input_topic = app.topic(INPUT_TOPIC, value_type=FunctionQuery)

builder = FunctionBuilder(app.logger)


@app.agent(input_topic)
async def handle_events(events):
    async for event in events:
        app.logger.info(f'Starting to build {event.function_id}')
        try:
            builder.process_request(event.function_id)
        except Exception as e:
            app.logger.error(f'Failed to build function {event.function_id}. Error: {repr(e)}')
        await asyncio.sleep(0.1)


async def start_worker(worker: faust.Worker) -> None:
    await worker.start()


if __name__ == '__main__':
    loop = asyncio.get_event_loop()
    worker = faust.Worker(app, loop=loop, loglevel=logging.INFO)
    try:
        loop.run_until_complete(start_worker(worker))
    finally:
        worker.stop_and_shutdown()
