from sqlmodel import create_engine, SQLModel

from models import Endpoint


DATABASE_URL = ""

engine = create_engine(DATABASE_URL, echo=True)

# Create the actual database using the engine
SQLModel.metadata.create_all(engine)