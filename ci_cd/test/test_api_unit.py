import pytest
import sys
import os

root = os.path.realpath(os.path.dirname(__file__) + "/../..")
services = root + "/services"

sys.path.append(os.path.realpath(root))
sys.path.append(os.path.realpath(services))

from services.api.routers.crud import crud
from services.api.routers.crud.crud import Location


@pytest.fixture
def zipcode():
    return 11206


def test_lat_long_to_zipcode(zipcode):
    coords = crud.zipcode_to_latlong(zipcode)
    assert isinstance(coords, Location)
