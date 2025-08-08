package middlewares

import (
	"context"
	"fmt"
	"go-cqrs-api/infrastructure"
	"net/http"
)

type contextDeviceApiKey string

const DeviceIDKey contextDeviceApiKey = "deviceId"

func DeviceIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db := infrastructure.GetDB()

		deviceIDHeader := r.Header.Get("x-api-key")
		if deviceIDHeader == "" {
			http.Error(w, "Missing API-KEY header", http.StatusBadRequest)
			return
		}

		var query string = `SELECT device_id FROM device_api_keys WHERE api_key = ?`

		var deviceID int64
		err := db.QueryRow(query, deviceIDHeader).Scan(&deviceID)
		if err != nil {
			fmt.Println("Failed to fetch device ID:", err)
			http.Error(w, "Invalid or missing API key", http.StatusUnauthorized)
			return
		}

		fmt.Println("Device ID from API key header:", deviceID)

		// Inject device ID into context
		ctx := context.WithValue(r.Context(), DeviceIDKey, deviceID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
