import uuid
from datetime import datetime, timezone

from sqlmodel import Field, Relationship, SQLModel


class Endpoint(SQLModel):
    id: uuid.UUID = Field(default_factory=uuid.uuid4, primary_key=True)
    url: str
    name: str | None
    is_up: bool
    is_monitored: bool
    status_code: int | None
    last_checked: datetime  # when it was pinged last
    created_at: datetime = Field(default_factory=datetime.now(timezone.utc))
    updated_at: datetime = Field(default_factory=datetime.now(timezone.utc))    # when it was updated in the db