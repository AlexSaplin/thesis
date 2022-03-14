import os


class Config(object):

    def __init__(self):
        self._kafka_broker = os.getenv('KAFKA_BROKER', 'kafka://localhost:9092')
        self._input_topic = os.getenv('INPUT_TOPIC', 'run_reports')
        self._gorilla_target = os.getenv('GORILLA_TARGET')

        self._billing_frequency_seconds = int(os.getenv('BILLING_FREQUENCY_SECONDS', '15'))

    @property
    def kafka_broker(self) -> str:
        return self._kafka_broker

    @property
    def input_topic(self) -> str:
        return self._input_topic

    @property
    def gorilla_target(self) -> int:
        return self._gorilla_target

    @property
    def billing_frequency_seconds(self) -> int:
        return self._billing_frequency_seconds
