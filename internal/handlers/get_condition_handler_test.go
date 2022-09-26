package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/zap"

	"github.com/kynetiv/simple-weather/internal/client"
	"github.com/kynetiv/simple-weather/internal/config"
	"github.com/kynetiv/simple-weather/internal/models"
)

const (
	v1ConditionsPath = "/api/v1/conditions"
	sfLatLonQuery = "?lat=37.6629&lon=-122.433"
)

func TestGetConditionHandler_ServeHTTP(t *testing.T) {

	// setup
	data := getValidPayload()
	resBytes, err := json.Marshal(data)
	if err != nil {
		t.Fatal("failed to marshal test data")
	}

	srv := getFakeServer(http.StatusOK, resBytes)

	cfg := &config.Config{
		APIKey:   "XXX",
		Endpoint: srv.URL + v1ConditionsPath,
	}

	ws := GetConditionHandler{
		Client: client.New(cfg),
		Logger: zap.NewExample(),
	}

	req := httptest.NewRequest(http.MethodGet, srv.URL+v1ConditionsPath+ sfLatLonQuery, nil)
	w := httptest.NewRecorder()

	// act
	ws.ServeHTTP(w, req)

	// validate
	res := w.Result()
	bodyBytes, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		t.Fatal("failed to read from body")
	}

	var got models.ConditionResponse
	err = json.Unmarshal(bodyBytes, &got)
	if err != nil {
		t.Fatal("failed to unmarshal response data")
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got: %d", http.StatusOK, res.StatusCode)
	}

	if got.FeelsLike != "Cold" {
		t.Errorf("expected feels_like: Cold, got: %s", got.FeelsLike)
	}
}

func getFakeServer(code int, payload []byte) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		w.Write(payload)
		return
	}))
}

func getValidPayload() models.OpenWeatherOneCallResponse {
	return models.OpenWeatherOneCallResponse{
		Lat:      37.6629,
		Lon:      -122.433,
		Timezone: "America/Los_Angeles",
		Current: models.CurrentConditions{
			Temp:      float64Ptr(57.4),
			FeelsLike: float64Ptr(58.4),
			Weather: []*models.WeatherItem{
				{
					ID:          802,
					Main:        "Clouds",
					Description: "scattered clouds",
					Icon:        "03n",
				},
			},
		},
		Alerts: nil,
	}
}

func float64Ptr(v float64) *float64 {
	return &v
}
