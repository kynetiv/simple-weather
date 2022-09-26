package main

import (
	"go.uber.org/zap"

	"github.com/kynetiv/simple-weather/internal"
	"github.com/kynetiv/simple-weather/internal/config"
	"github.com/kynetiv/simple-weather/internal/server"
)

func main() {
	logger, _ := zap.NewProduction()
	logger.Info("init weather-service-api")

	cfg, err := config.New()
	if err != nil {
		logger.Fatal("failed to load config", zap.Error(err))
	}

	api := internal.NewWeatherService(logger, cfg)

	srv := server.New(api)

	srv.Start()
}
