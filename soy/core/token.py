import requests
import os

def get_token():
    API_KEY = os.getenv("API_KEY")
    API_SECRET = os.getenv("API_SECRET")
    if API_KEY is None or API_SECRET is None:
        raise Exception("API_KEY and API_SECRET are not valid environment variables")

    response = requests.post("https://api.twitter.com/oauth2/token?grant_type=client_credentials", auth=(API_KEY, API_SECRET))

    res_json = response.json() 
    if "access_token" in res_json:
        return res_json["access_token"]
    else:
        raise Exception(f"unexpected response from twitter's oauth {res_json}")
