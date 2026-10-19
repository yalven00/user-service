package config

import (
	"os"
	"time"
)

type Config struct {
	DBUser          string
	DBPassword      string
	DBHost          string
	DBPort          string
	DBName          string
	ServerPort      string
	JWTSecret       string
	JWTAccessExpiry time.Duration
}

func Load() *Config {
	return &Config{
		DBUser:          getEnv("DB_USER", "root"),
		DBPassword:      getEnv("DB_PASSWORD", "12345!"),
		DBHost:          getEnv("DB_HOST", "127.0.0.1"),
		DBPort:          getEnv("DB_PORT", "3306"),
		DBName:          getEnv("DB_NAME", "userdb"),
		ServerPort:      getEnv("SERVER_PORT", "8080"),
		JWTSecret:       getEnv("JWT_SECRET", "change-me-in-production"),
		JWTAccessExpiry: 24 * time.Hour,
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}