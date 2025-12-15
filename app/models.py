from datetime import datetime
import uuid

from sqlmodel import Field, Relationship, SQLModel

class Endpoint(SQLModel):
    id: uuid.UUID
    url: str
    is_active: bool
    last_checked: str