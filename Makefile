SHELL := /bin/bash

API_KEY ?="XXXX"
ENDPOINT ?="https://api.openweathermap.org/data/3.0"

default: run

run:
	@OPEN_WEATHER_API_KEY=${API_KEY} \
	OPEN_WEATHER_URL=${ENDPOINT} \
	go run main.go


miami:
	curl -sk "http://localhost:8080/api/v1/conditions?lat=25.7617&lon=-80.1918"

dallas:
	curl -sk "http://localhost:8080/api/v1/conditions?lat=32.779167&lon=-96.808891"

sf:
	curl -sk "http://localhost:8080/api/v1/conditions?lat=37.662937&lon=-122.433014"