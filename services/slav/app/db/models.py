from uuid import uuid4

from sqlalchemy import Column, Integer, String
from sqlalchemy.dialects.postgresql import UUID, JSONB
from sqlalchemy.ext.declarative import declarative_base

BaseCls = declarative_base()


class ContainerMeta(BaseCls):
    __tablename__ = 'containers'

    id = Column(UUID(as_uuid=True), primary_key=True, default=uuid4)
    name = Column(String, nullable=False, unique=True)
    scale = Column(Integer, nullable=False)
    instance_type = Column(String, nullable=False)
    image = Column(String, nullable=False)
    port = Column(Integer, nullable=False)
    owner_id = Column(String, nullable=False)
    env = Column(JSONB, nullable=True)
    auth = Column(JSONB, nullable=True)

    @classmethod
    def from_dict(cls, d: dict):
        return cls(**d)
