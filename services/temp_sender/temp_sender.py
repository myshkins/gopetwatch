# from sense_hat import SenseHat
import requests
from datetime import datetime as dt
import json
import logging


logging.basicConfig(
    filename='temp_sender.log', encoding='utf-8', level=logging.INFO
    )

# hat = SenseHat()


def send_temp_point(arg):
    # temp = hat.get_temperature()
    temp = 70.5
    now = dt.now
    data_dict = {"temperature": temp, "time": now}
    j = json.dumps(data_dict)
    response = requests.post("http://0.0.0.0:8081/post", json=j)
    logging.info("time: %s, response: %s", now, response)
