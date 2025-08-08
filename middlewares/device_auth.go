package middlewares

import (
	"context"
	"net/http"
	"strconv"
)

type contextDeviceApiKey string

const DeviceIDKey contextDeviceApiKey = "deviceId"

func DeviceIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		deviceIDHeader := r.Header.Get("API-KEY")
		if deviceIDHeader == "" {
			http.Error(w, "Missing API-KEY header", http.StatusBadRequest)
			return
		}

		deviceID, err := strconv.ParseInt(deviceIDHeader, 10, 64)
		if err != nil {
			http.Error(w, "Invalid API-KEY header", http.StatusBadRequest)
			return
		}

		// Inject device ID into context
		ctx := context.WithValue(r.Context(), DeviceIDKey, deviceID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
