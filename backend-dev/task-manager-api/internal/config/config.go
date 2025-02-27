package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        int
	Environment string
	DatabaseURL string
	LogLevel    string
}

func Load() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	port, err := strconv.Atoi(getEnv("PORT", "8080"))
	if err != nil {
		return nil, fmt.Errorf("invalid port: %w", err)
	}

	return &Config{
		Port:        port,
		Environment: getEnv("ENVIRONMENT", "development"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:postgres@172.22.0.2:5432/task_db?sslmode=disable"),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
