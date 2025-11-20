package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	MneeEnv    string
	MneeApiKey string
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	return &Config{
		Port:       getEnv("PORT", "8080"),
		MneeEnv:    getEnv("MNEE_ENV", "sandbox"),
		MneeApiKey: getEnv("MNEE_API_KEY", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
