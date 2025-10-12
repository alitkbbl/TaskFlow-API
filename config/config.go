package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DBPath        string
	Port          string
	JWTSecret     string // secret key
	JWTExpMinutes int    // expire
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

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
