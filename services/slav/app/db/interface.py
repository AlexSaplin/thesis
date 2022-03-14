from contextlib import contextmanager

from sqlalchemy import create_engine
from sqlalchemy.exc import IntegrityError
from sqlalchemy.orm import sessionmaker, Session

from errors import ContainerAlreadyExistsException
from slav_dataclasses import InstanceType
from db.adapter import SlavAdapter
from db.models import ContainerMeta
from tesseract.client import TesseractClient


class DatabaseInterface:

    def __init__(self, db_url: str, tesseract_client: TesseractClient, logger):
        self.session_maker = sessionmaker(bind=create_engine(db_url))
        self.tesseract_client = tesseract_client
        self.logger = logger

    @contextmanager
    def _connect(self):
        session: Session = self.session_maker()
        try:
            yield session
            session.commit()
        except Exception as e:
            session.rollback()
            self.logger.error(f'Transaction failed with error: {e}')
            raise
        finally:
            session.close()

    def create_container(self,
                         owner_id: str,
                         name: str,
                         scale: int,
                         instance_type: InstanceType,
                         image: str,
                         port: int,
                         env: list,
                         auth: dict):
        try:
            with self._connect() as session:
                new_container = ContainerMeta(
                    name=name,
                    scale=scale,
                    instance_type=instance_type.value,
                    image=image,
                    port=port,
                    owner_id=owner_id,
                    env=env,
                    auth=auth,
                )
                session.add(new_container)
                session.commit()
                self.logger.info(f'Container in db has been created. Name: {name}, OwnerID: {owner_id}')
                self.tesseract_client.apply_changes(name, owner_id, scale, port, image, instance_type, env, auth)
                return SlavAdapter.container_meta_to_dataclass(new_container)
        except IntegrityError:  # Handle unique exception
            self.logger.error(f'Container with name {name} already exist')
            raise ContainerAlreadyExistsException(f'Container with name {name} already exist')
        except Exception as e:
            self.logger.error(f'Container creation in DB failed error: {e}')
            raise

    def update_container(self, owner_id: str, name: str, params: dict):
        try:
            with self._connect() as session:
                container = session.query(ContainerMeta).filter(ContainerMeta.name == name,
                                                                ContainerMeta.owner_id == owner_id).first()
                for key, value in params.items():
                    setattr(container, key, value)

                self.tesseract_client.apply_changes(name, owner_id, container.scale, container.port, container.image,
                                                    InstanceType(container.instance_type), container.env,
                                                    container.auth)
                return SlavAdapter.container_meta_to_dataclass(container)
        except Exception as e:
            self.logger.error(f'Container updating in DB failed. Error: {e}')
            raise

    def delete_container(self, owner_id: str, name: str):
        try:
            with self._connect() as session:
                container = session.query(ContainerMeta).filter(ContainerMeta.name == name,
                                                                ContainerMeta.owner_id == owner_id).first()
                session.delete(container)

                self.tesseract_client.delete_container(name, owner_id)
        except Exception as e:
            self.logger.error(f'Container deleting in DB failed. Error: {e}')
            raise

    def get_container(self, owner_id: str, name: str):
        try:
            with self._connect() as session:
                status, error = self.tesseract_client.get_status(name, owner_id)

                container = session.query(ContainerMeta).filter(ContainerMeta.name == name,
                                                                ContainerMeta.owner_id == owner_id).first()
                return SlavAdapter.container_meta_to_dataclass(container), status, error
        except Exception as e:
            self.logger.error(f'Container fetching in DB failed. Error: {e}')
            raise

    def list_containers(self, owner_id: str):
        try:
            with self._connect() as session:
                containers = session.query(ContainerMeta).filter(ContainerMeta.owner_id == owner_id).all()
                return [SlavAdapter.container_meta_to_dataclass(container) for container in containers]
        except Exception as e:
            self.logger.error(f'Failed to fetch containers for user {owner_id}. Error: {e}')
            raise

    def list_all_containers(self):
        try:
            with self._connect() as session:
                containers = session.query(ContainerMeta).all()
                return [SlavAdapter.container_meta_to_dataclass(container) for container in containers]
        except Exception as e:
            self.logger.error(f'Failed to fetch all containers. Error: {e}')
            raise
