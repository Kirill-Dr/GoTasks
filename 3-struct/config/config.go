package config

import (
	"os"
)

type Config struct {
	Key string
}

func NewConfig() *Config {
	key := os.Getenv("KEY")
	if key == "" {
		panic("Key is not set in .env file")
	}
	return &Config{
		Key: key,
	}
}
