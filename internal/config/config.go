package config

import "github.com/kelseyhightower/envconfig"

// Config holds application configuration values.
type Config struct {
	Port     string `envconfig:"PORT" default:":8080"`                  // Server port
	SWAPIURL string `envconfig:"SWAPI_URL" default:"https://swapi.dev"` // SWAPI base URL
}

// Load reads environment variables and returns a Config instance.
func Load() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
