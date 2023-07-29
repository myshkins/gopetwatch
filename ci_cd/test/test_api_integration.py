import pytest
import requests
import sys
import os
from sqlalchemy import create_engine


root = os.path.realpath(os.path.dirname(__file__) + "/../..")
api_path = root + "/services/api/"

sys.path.append(os.path.realpath(root))
sys.path.append(os.path.realpath(api_path))


ZIPCODES = ["11206", 80304, "01913", 99723, "96712"]
ZIPCODES_INVALID = ["Brooklyn, NY", 112060, "11206-1839", "00000", "-----"]
POSTGRES_URI = os.environ["POSTGRES_URI"]
STATION_IDS = ["840360470118", "840360610134", "840360810120"]


@pytest.fixture
def invalid_station_ids():
    return ["not and id", "", 0, 1, "0000", 000000000000, "000000000000"]


@pytest.fixture
def station_ids():
    return ["840360470118", "840360610134", "840360810120"]


@pytest.fixture
def engine():
    engine = create_engine(POSTGRES_URI)
    return engine


def test_health_check():
    """
    GIVEN   containers are running
    WHEN    health check endpoint is called with GET method
    THEN    response with status 200 and body OK is returned
    """
    response = requests.get("http://airboo_api:10100/health-check")
    assert response.status_code == 200
    assert response.json() == {"message": "OK"}


@pytest.mark.parametrize("zipcode", ZIPCODES)
def test_pos_get_nearby_stations(zipcode):
    """
    GIVEN   valid zipcode argument
    WHEN    get_nearby_stations endpoint is called
    THEN    response status is 200 and
            response body is a list of five stations that conform to the
            StationsAirnowPydantic model
    """
    response = requests.get(
        f"http://airboo_api:10100/api/stations/all-nearby/?zipcode={zipcode}"
    )
    assert response.status_code == 200
    station_list = response.json()
    assert len(station_list) == 5
    assert all(station_list)


@pytest.mark.parametrize("zipcode", ZIPCODES_INVALID)
def test_neg_get_nearby_stations(zipcode):
    """
    GIVEN   an invalid zipcode argument
    WHEN    get_nearby_stations endpoint is called
    THEN    response body is 400 and response body is error message
    """
    response = requests.get(
        f"http://airboo_api:10100/api/stations/all-nearby/?zipcode={zipcode}"
    )
    assert response.status_code == 400
    assert response.json() == {"detail": "invalid zipcode"}


@pytest.mark.parametrize("period,length", [("twelve_hr", 12), ("twenty_four_hr", 24)])
def test_pos_get_readings_from_ids(station_ids, period, length):
    """
    GIVEN   valid station id, time period, and pollutant arguments
    WHEN    get_readings_from_ids endpoint is called
    THEN    response status is 200 and
            response body is list of readings conforming to ReadingsAirnowPydantic model
            and length of list corresponds to the time period argument passed
    """
    id_lst = [f"?ids={id}&" for id in station_ids]
    id_str = "".join(id_lst)
    query = id_str + f"period={period}"
    response = requests.get(
        f"http://airboo_api:10100/api/air-readings/from-ids/{query}"
    )
    assert response.status_code == 200
    assert len(response.json()[0]["readings"]) == length


def test_neg_get_readings_from_ids():
    """
    GIVEN   invalid station id arguments and valid time period argument
    WHEN    get_readings_from_ids enpoint is called
    THEN    response status is 400 and response body is error message
    """
    pass


def test_get_pollutants(station_ids):
    """
    given   valid station ids
    when    get_pollutants endpoint is called
    then    response status is 200 and
            response body is list of pollutants
    """
    id_lst = [f"?ids={id}&" for id in station_ids]
    id_str = "".join(id_lst)
    query = id_str.removesuffix("&")
    response = requests.get(
        f"http://airboo_api:10100/api/air-readings/from-ids/{query}"
    )
    pollutants = response.json()
    assert response.status_code == 200
    assert 1 >= len(pollutants) <= 6
