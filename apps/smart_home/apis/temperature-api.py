import random
from fastapi import FastAPI
from datetime import datetime
#pyth0n 3.8
import pytz

app = FastAPI()

@app.get("/temperature")
def temperature_bylocation(location: str = "", sensorID: str =""):
    # If no location is provided, use a default based on sensor ID
    if location == "":
        if sensorID == "1": 
            location = "Living Room"
        elif  sensorID == "2":
            location = "Bedroom"
        elif sensorID == "3":
            location = "Kitchen"
        else:
            location = "Unknown"
	

	# If no sensor ID is provided, generate one based on location
    if sensorID == "" :
        if location == "Living Room":
            sensorID = "1"
        elif location == "Bedroom":
            sensorID = "2"
        elif location == "Kitchen":
            sensorID = "3"
        else:
            sensorID = "0"
    
    now = datetime.now(pytz.timezone('Europe/Moscow'))
    return {
		"value":       random.randint (-100, 100),
        "unit":        "Celcius",
        "timestamp":   now,
        "location":    location,
		"status":      "active",
        "sensor_id":   str(sensorID),
		"sensor_type": "Temperature",
		"description": "Temperature by location",
	}

@app.get("/temperature/{sensor_id}")
def temperature_byID(sensor_id: int):
    now = datetime.now(pytz.timezone('Europe/Moscow'))
    return {
        "value":       random.randint (-100, 100),
        "unit":        "Celcius",
        "timestamp":   now,
        "location":    " ",
		"status":      "healthy",
        "sensor_id":   str(sensor_id),
		"sensor_type": "Temperature",
		"description": "Temperature by sensor ID",
	}
    
    