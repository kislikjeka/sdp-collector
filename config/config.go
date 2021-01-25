package config

import "os"

type Config struct {
	cf CollectorConfig
}

type CollectorConfig struct {
	name string
}

func NewConfig() *Config {
	collectorName := os.Getenv("CollectorName")
	return &Config{CollectorConfig{name: collectorName}}
}

func (c *CollectorConfig) Name() string {
	return c.name
}
