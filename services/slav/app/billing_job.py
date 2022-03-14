from datetime import datetime

from db.adapter import SlavAdapter
from db.interface import DatabaseInterface
from gorilla.client import GorillaClient
from slav_dataclasses import StateType
from tesseract.client import TesseractClient


def job(db_interface: DatabaseInterface, tesseract_client: TesseractClient, gorilla_client: GorillaClient, logger):
    try:
        logger.info(f'STARTING SLAV BILLING JOB AT {str(datetime.utcnow())}')
        containers = db_interface.list_all_containers()
        deltas = []
        today_dt = datetime.utcnow()
        today_dt = today_dt.replace(hour=0, minute=0, second=0, microsecond=0)
        for container in containers:
            state, _ = tesseract_client.get_status(container.Name, container.OwnerID)
            if state == StateType.RUNNING:
                new_balance = -SlavAdapter.instance_type_to_billing_price(container.Instance).value * container.Scale
                deltas.append(
                    {
                        "date": int(today_dt.timestamp()),
                        "category": "CONTAINERS",
                        "balance": new_balance,
                        "owner_id": container.OwnerID,
                        "object_id": container.ID,
                        "object_type": "CONTAINER",
                    }
                )
        if deltas:
            gorilla_client.add_deltas(deltas)
    except Exception as e:
        logger.error(f'Execution of billing job has been failed at {str(datetime.utcnow())}. Error: {repr(e)}')
