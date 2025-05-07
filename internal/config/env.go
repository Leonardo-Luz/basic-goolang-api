package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	POSTGRES_HOST     string
	POSTGRES_PORT     string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DATABASE string
}

var Envs = initConfig()

func initConfig() Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env")
	}

	return Config{
		POSTGRES_HOST:     getEnv("DEV_POSTGRES_HOST", "localhost"),
		POSTGRES_PORT:     getEnv("DEV_POSTGRES_PORT", "5432"),
		POSTGRES_USER:     getEnv("DEV_POSTGRES_USER", "postgres"),
		POSTGRES_PASSWORD: getEnv("DEV_POSTGRES_PASSWORD", "postgres"),
		POSTGRES_DATABASE: getEnv("DEV_POSTGRES_DATABASE", "database"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
