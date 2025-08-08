package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"go-cqrs-api/infrastructure"
	"time"
)

func GenerateAndStoreAPIKey(deviceID int64) (string, error) {
	db := infrastructure.GetDB()
	// Generate a secure 32-byte API key
	apiKeyBytes := make([]byte, 32)
	if _, err := rand.Read(apiKeyBytes); err != nil {
		return "", fmt.Errorf("failed to generate API key: %w", err)
	}
	apiKey := hex.EncodeToString(apiKeyBytes)

	// Insert the API key into the database
	query := `
		INSERT INTO device_api_keys (device_id, api_key, created_at)
		VALUES (?, ?, ?)
	`

	_, err := db.Exec(query, deviceID, apiKey, time.Now())
	if err != nil {
		return "", fmt.Errorf("failed to insert API key into DB: %w", err)
	}

	return apiKey, nil
}
