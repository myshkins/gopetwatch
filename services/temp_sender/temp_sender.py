# temp sender script, runs every 5 minutes via crontab
import logging
import os
import random
import time
import requests
from dotenv import load_dotenv
from requests.auth import HTTPBasicAuth
from datetime import timezone, datetime as dt


environment = "dev"
load_dotenv()
if os.environ.get("LOGNAME") == "iceking":
    from sense_hat import SenseHat
    environment = "prod"

logging.basicConfig(
    filename='temp_sender.log', encoding='utf-8', level=logging.INFO
)

def send_temp_point():
    now = str(dt.now(tz=timezone.utc))
    if environment == "dev":
        temp = random.uniform(60.0, 70.0)
        url = "http://0.0.0.0:8081/post"
    else:
        hat = SenseHat()
        t = (9/5) * hat.get_temperature() + 32
        temp = round(t, 2)
        if type(temp) != float:
            logging.warn("hat.get_temperature returned non-float value: %s, temp")
        url = "https://gopetwatch.ak0.io/post"
    data_dict = {"temperature": round(temp, 2), "datetime": now}
    basic = HTTPBasicAuth(os.environ.get("GIN_USER_ID"), os.environ.get("GIN_PW"))
    response = requests.post(url=url, json=data_dict, auth=basic)
    response_code = response.status_code
    if response_code != 200:
        logging.warning(
            "post at time: {}, got response {}".format(now, response_code)
        )

if __name__ == "__main__":
    send_temp_point()
