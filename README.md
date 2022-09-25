# Simple Weather API

### Overview

This simple-weather project is a http server that provides an API to get the current weather conditions data from the [Open Weather Map API](https://openweathermap.org/api). Making a request with any valid latitude and longitude geo coordinates returns the current weather conditions in that area in JSON format.

### Pre-reqs
1. Recent version of Go (Golang) installed. [Download](https://go.dev/doc/install)
1. Need an API KEY [Sign-up](https://home.openweathermap.org/users/sign_up) and Subscribe for the Free [One-Call](https://openweathermap.org/api/one-call-3) service.

### Usage
1. Start the Server with your API KEY
```sh
API_KEY=XXXX make run
```

2. Find your lat, lon geo coordinates.

3. Curl the weather-service API (See Makefile for examples)
```sh
curl -sk "http://localhost:8080/api/v1/conditions?lat=25.7617&lon=-80.1918"
```
   
Examples:

Miami
```sh
make miami
{
  "alerts": [
    {
      "sender_name": "NWS Miami (Southern Florida)",
      "event": "Coastal Flood Statement",
      "start": 1664104440,
      "end": 1664172000,
      "description": "...COASTAL FLOOD STATEMENT IN EFFECT THROUGH LATE TONIGHT...\n* WHAT...Isolated minor coastal flooding.\n* WHERE...Coastal Broward and Coastal Miami-Dade Counties.\n* WHEN...Through late tonight.\n* IMPACTS...Some water on low lying roads and property during high\ntide.",
      "tags": [
        "Coastal event",
        "Flood"
      ]
    }
  ],
  "summary": "Current conditions: Clouds (broken clouds) and 1 alert(s) in the area.",
  "temp": 90.25,
  "feels_temp": 102.85,
  "feels_like": "Blazing Hot"
}
```

Dallas

```sh
make dallas
{
  "summary": "Current conditions: Clouds (scattered clouds) and 0 alert(s) in the area.",
  "temp": 95.23,
  "feels_temp": 102.4,
  "feels_like": "Blazing Hot"
}
```

San Francisco

```sh
make sf
{
  "summary": "Current conditions: Clouds (scattered clouds) and 0 alert(s) in the area.",
  "temp": 65.71,
  "feels_temp": 65.84,
  "feels_like": "Cool"
}
```



