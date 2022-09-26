package internal

import (
	"go.uber.org/zap"

	"github.com/kynetiv/simple-weather/internal/client"
	"github.com/kynetiv/simple-weather/internal/config"
	"github.com/kynetiv/simple-weather/internal/handlers"
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
