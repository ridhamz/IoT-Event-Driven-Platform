package config

import (
	"log"
	"os"
)

// Config holds the application configuration
type Config struct {
	RabbitURL string
	RedisAddr string
	DBPath    string
}

// AppConfig is the global configuration instance
var AppConfig Config

// Load reads environment variables and sets the configuration
func Load() {
	AppConfig = Config{
		RabbitURL: getEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/"),
		RedisAddr: getEnv("REDIS_ADDR", "localhost:6379"),
		DBPath:    getEnv("DB_PATH", "./data.db"),
	}
}

// getEnv returns the value of an env variable or fallback if not set
func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Printf("Using default for %s: %s", key, fallback)
		return fallback
	}
	return val
}
