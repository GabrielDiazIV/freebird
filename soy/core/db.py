import psycopg2
import os
from psycopg2 import OperationalError
from psycopg2.extras import RealDictCursor

def create_connection():
    connection = None
    try:
        connection = psycopg2.connect(
            host=os.getenv("POSTGRES_HOST"),
            database=os.getenv("POSTGRES_NAME"),
            user=os.getenv("POSTGRES_USER"),
            password=os.getenv("POSTGRES_PASS"),
            cursor_factory=RealDictCursor
        )
        print("Connected to DB successfully")
    except OperationalError as e:
        print("DB connection error: " + e)
    return connection
def query(query_str,connection):
    cursor = connection.cursor()
    try:
        cursor.execute(query_str)
        return cursor.fetchall()
    except OperationalError as e:
        print(f"DB error: '{e}'")
def update(update_str,connection):
    cursor = connection.cursor()
    try:
        cursor.execute(update_str)
        connection.commit()
    except OperationalError as e:
        print(f"DB error: '{e}'")
def disconnect(connection):
    try:
        connection.close()
        print("DB connection terminated successfully")
    except OperationalError as e:
        print(f"DB disconnection error: '{e}'")
