package config

import (
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Port     string    `mapstructure:"PORT"`
	LogLevel log.Level `mapstructure:"LOG_LEVEL"`
}

type Provider interface {
	Load() (*Config, error)
}
