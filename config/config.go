package config

import (
	"os"
)

type Config struct {
	DBPath        string
	Port          string
	JWTSecret     string
	JWTExpMinutes int
}

func LoadConfig() *Config {
	return &Config{
		DBPath:        getEnv("DB_PATH", "taskflow.db"),
		Port:          getEnv("PORT", "8080"),
		JWTSecret:     getEnv("JWT_SECRET", "supersecret"),
		JWTExpMinutes: 60,
	}
}

func getEnv(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}
