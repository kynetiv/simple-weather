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
	cfgExclude  = "OPEN_WEATHER_EXCLUDE"
	cfgUnit     = "OPEN_WEATHER_UNIT"
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

	var ok bool
	if cfg.Exclude, ok = os.LookupEnv(cfgExclude); !ok {
		cfg.Exclude = "minutely,hourly,daily" // default
	}

	if cfg.Exclude, ok = os.LookupEnv(cfgUnit); !ok {
		cfg.Unit = "imperial" // default
	}

	return cfg, nil
}
