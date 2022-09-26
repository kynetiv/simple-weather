package config

import (
	"fmt"
	"os"
)

type Config struct {
	APIKey   string
	Endpoint string
	Exclude  string
	Unit     string
}

const (
	cfgAPIKey   = "OPEN_WEATHER_API_KEY"
	cfgEndpoint = "OPEN_WEATHER_URL"
)

func New() (*Config, error) {
	cfg := &Config{}

	cfg.APIKey = os.Getenv(cfgAPIKey)
	if cfg.APIKey == "" {
		return cfg, fmt.Errorf("bad config: missing %s", cfgAPIKey)
	}

	cfg.Endpoint = os.Getenv(cfgEndpoint)
	if cfg.Endpoint == "" {
		return cfg, fmt.Errorf("bad config: missing %s", cfgEndpoint)
	}

	cfg.Exclude = "minutely,hourly,daily" // default
	cfg.Unit = "imperial" // default

	return cfg, nil
}
