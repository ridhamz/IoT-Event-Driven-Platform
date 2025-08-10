package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	SQS_URL    string
	JWT_SECRET string
	DB_URL     string
	S3_BUCKET  string
}

// AppConfig is the global configuration instance
var AppConfig Config

// Load reads environment variables and sets the configuration
func Load() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	AppConfig = Config{
		SQS_URL:    getEnv("SQS_URL"),
		DB_URL:     getEnv("DB_URL"),
		JWT_SECRET: getEnv("JWT_SECRET"),
		S3_BUCKET:  getEnv("S3_BUCKET"),
	}
}

// getEnv returns the value of an env variable or fallback if not set
func getEnv(key string) string {
	val := os.Getenv(key)
	fmt.Printf("Env %s: %s\n", key, val)
	if val == "" {
		log.Printf("Env is empty %s", key)
		os.Exit(1)
	}
	return val
}
