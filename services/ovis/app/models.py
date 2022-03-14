import uuid
import pydantic


class Message(pydantic.BaseModel):
    model_id: uuid.UUID
