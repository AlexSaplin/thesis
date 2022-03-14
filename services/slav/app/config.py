from enum import Enum
import os


DB_URL = 'postgres://postgres:postgres@postgres_slav:5432/postgres?sslmode=disable'
SLAV_PORT = 8086
TESSERACT_TARGET = os.getenv('TESSERACT_TARGET')
GORILLA_TARGET = os.getenv('GORILLA_TARGET')

CONTAINERS_DOMAIN = 'deepmux-containers.com'


class BillingPrice(Enum):
    STARTER_PRICE = 0.02 / 60
    INFERENCE_PRICE = 0.31 / 60
