package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"go.uber.org/zap"

	"simple-weather/internal/client"
	"simple-weather/internal/models"
)

type GetConditionHandler struct {
	Client *client.OpenWeatherClient
	Logger *zap.Logger
}

func (h *GetConditionHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	logger := h.Logger.Named("get-condition-handler")
	logger.Info("Request received")

	params := req.URL.Query()
	enc := json.NewEncoder(w)

	res, err := h.getCondition(h.Logger, params)
	if err != nil {
		logger.Error("failed to get condition", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		err = enc.Encode("Server Error")
		if err != nil {
			logger.Error("failed to encode json response")
		}
		return
	}

	err = enc.Encode(res)
	if err != nil {
		logger.Error("failed to encode json response")
	}
	logger.Info("Success", zap.Any("Response", res))

	return
}

func (h *GetConditionHandler) getCondition(logger *zap.Logger, p url.Values) (resp models.ConditionResponse, err error) {
	logger.Info("get-condition", zap.Any("params", p))

	var (
		lat, lon string
	)

	if lat = p.Get("lat"); lat == "" {
		return resp, fmt.Errorf("lat query param is required")
	}

	if lon = p.Get("lon"); lon == "" {
		resp.Error = "lon query param is required"
		return resp, fmt.Errorf("bad request")
	}

	res, err := h.Client.GetConditions(lat, lon)
	if err != nil {
		return resp, fmt.Errorf("do request failed : %w", err)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	var openData models.OpenWeatherOneCallResponse

	err = json.Unmarshal(body, &openData)
	if err != nil {
		logger.Error("raw weather response", zap.Any("body", string(body)))
		return resp, fmt.Errorf("failed to unmarshal open weather data: %w", err)
	}

	if openData.Code != nil {
		rawRes := string(body)
		logger.Error("bad response from open api", zap.Any("body", rawRes))
		return resp, fmt.Errorf("request failed", rawRes)
	}

	logger.Debug("weather data", zap.Any("raw", openData))

	return h.getConditionResponse(openData), nil
}

func (h *GetConditionHandler) getConditionResponse(data models.OpenWeatherOneCallResponse) models.ConditionResponse {
	return models.ConditionResponse{
		Alerts:    data.Alerts,
		Summary:   h.getSummaryDescription(*data.Current.Weather[0], len(data.Alerts)),
		Temp:      data.Current.Temp,
		FeelsTemp: data.Current.FeelsLike,
		FeelsLike: h.getTempDescription(*data.Current.FeelsLike),
	}
}

func (h *GetConditionHandler) getSummaryDescription(weather models.WeatherItem, alertCount int) string {
	return fmt.Sprintf("Current conditions: %s (%s) and %d alert(s) in the area.", weather.Main, weather.Description, alertCount)
}

func (h *GetConditionHandler) getTempDescription(feelsTemp float64) string {
	var tempDesc string
	if feelsTemp <= float64(32) {
		tempDesc = "Freezing Cold"
	} else if feelsTemp < float64(60) {
		tempDesc = "Cold"
	} else if feelsTemp < float64(67) {
		tempDesc = "Cool"
	} else if feelsTemp < float64(75) {
		tempDesc = "Moderate"
	} else if feelsTemp < float64(85) {
		tempDesc = "Warm"
	} else if feelsTemp < float64(90) {
		tempDesc = "Hot"
	} else if feelsTemp < float64(97) {
		tempDesc = "Hot"
	} else {
		tempDesc = "Blazing Hot"
	}

	return tempDesc
}
