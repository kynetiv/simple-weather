package internal

import (
	"go.uber.org/zap"

	"simple-weather/internal/client"
	"simple-weather/internal/config"
	"simple-weather/internal/handlers"
)

const (
	openWeatherMapURL = "https://api.openweathermap.org/data/2.5"
)

type WeatherService struct {
	GetCondition *handlers.GetConditionHandler
}

func NewWeatherService(logger *zap.Logger, cfg *config.Config) WeatherService {
	c := client.New(cfg)
	return WeatherService{
		GetCondition: &handlers.GetConditionHandler{Logger: logger, Client: c},
	}
}
