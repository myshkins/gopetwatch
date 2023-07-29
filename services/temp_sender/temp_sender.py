# from sense_hat import SenseHat
import requests
from datetime import datetime as dt
# import json
import logging
import random


logging.basicConfig(
    filename='temp_sender.log', encoding='utf-8', level=logging.INFO
    )

# hat = SenseHat()


def send_temp_point():
    now = str(dt.now())
    # temp = round(hat.get_temperature(), 2)
    # if type(temp) != float:
    #   logging.warn("hat.get_temperature returned non-float value: %s, temp")
    temp = random.uniform(60.0, 70.0)
    data_dict = {"temperature": round(temp, 2), "datetime": now}
    response = requests.post(url="http://0.0.0.0:8081/post", json=data_dict)
    response_code = response.status_code
    if response_code != 200:
        logging.warning(
            "post at time: {}, got response {}".format(now, response_code)
        )


send_temp_point()
