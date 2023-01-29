from psycopg2 import *
import os
from psycopg2 import OperationalError

connection = None

def create_connection():
    try:
        connection = connect()
        print("Connected to DB successfully")
    except OperationalError as e:
        print("DB connection error: " + e)
    return connection
def query(query_str):
    connection.autocommit = True
    cursor = connection.cursor()
    try:
        cursor.execute(query_str)
        return cursor.fetchall()
    except OperationalError as e:
        print(f"DB error: '{e}'")
def update(update_str):
    connection.autocommit = True
    cursor = connection.cursor()
    try:
        cursor.execute(update_str)
    except OperationalError as e:
        print(f"DB error: '{e}'")
def disconnect():
    try:
        connection.close()
        print("DB connection terminated successfully")
    except OperationalError as e:
        print(f"DB disconnection error: '{e}'")
