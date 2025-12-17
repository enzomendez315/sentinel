from sqlmodel import create_engine, SQLModel

from models import Endpoint
from config import settings

engine = create_engine(settings.DATABASE_URL, echo=True)

# Create the actual database using the engine
SQLModel.metadata.create_all(engine)