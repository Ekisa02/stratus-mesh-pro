package config

import (
	"os"
)

type Config struct {
	Port string
	DatabaseDSN string
}

func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	dsn := os.Getenv("DATABASE_DSN")
	return &Config{
		Port: port,
		DatabaseDSN: dsn,
	}
}
