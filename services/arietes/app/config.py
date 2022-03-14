import os


IBIS_TARGET = os.getenv('IBIS_TARGET')

REGISTRY_URL = os.getenv('REGISTRY_URL')

KAFKA_BROKER = os.getenv('KAFKA_BROKER')
INPUT_TOPIC = os.getenv('INPUT_TOPIC')

S3_TARGET = os.getenv('S3_TARGET')

AVAILABLE_IMAGES = [
    'python3.6',
    'python3.7',
    'python3.6-tensorflow2.1-pytorch-1.6-cuda10.1',
    'python3.7-tensorflow2.1-pytorch-1.6-cuda10.1',
    'python3.7-tensorflow2.2-pytorch-1.6-cuda10.1',
    'python3.7-tensorflow1.13.1-pytorch-1.3-cuda10.0',
    'python3.7-mmdetection-pytorch-1.6-cuda10.1',
]
