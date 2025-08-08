package api

import (
	"encoding/json"
	"fmt"
	"go-cqrs-api/commands"
	"go-cqrs-api/domain"
	"go-cqrs-api/events"
	"go-cqrs-api/middlewares"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func SetupDevicesRoutes() {
	router := GetRouter()

	router.Group(func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware)

		r.Post("/api/devices/register", func(w http.ResponseWriter, r *http.Request) {
			var body domain.Device
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}

			userIdCtx := r.Context().Value(middlewares.UserIDKey)
			userIDFloat, ok := userIdCtx.(float64)
			if !ok {
				http.Error(w, "User ID not found in context", http.StatusUnauthorized)
				return
			}

			userId := int64(userIDFloat) // Convert float64 to int64 if necessary

			// Here you can do validation or preprocessing on the device data if needed.

			device := domain.Device{
				ID:        body.ID, // ID can be omitted if auto-incremented
				Name:      body.Name,
				UserID:    userId,
				CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			}

			fmt.Println("Device to be created:", device)

			if err := commands.HandleCreateDevice(device); err != nil {
				fmt.Println("Error creating device:", err)
				http.Error(w, "Could not create device", http.StatusInternalServerError)
				return
			}

			response := map[string]string{
				"message": "Device created successfully",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})

		r.Get("/api/devices", func(w http.ResponseWriter, r *http.Request) {
			userIdCtx := r.Context().Value(middlewares.UserIDKey)
			userIDFloat, ok := userIdCtx.(float64)
			if !ok {
				http.Error(w, "User ID not found in context", http.StatusUnauthorized)
				return
			}

			userId := int64(userIDFloat)

			// Fetch user devices using command
			devices, err := commands.GetUserDevices(userId)
			if err != nil {
				http.Error(w, "Failed to retrieve user devices", http.StatusInternalServerError)
				return
			}

			// Respond with JSON
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(devices)

		})
	})

	router.Group(func(r chi.Router) {
		r.Use(middlewares.DeviceIDMiddleware)
		r.Post("/api/devices/data", func(w http.ResponseWriter, r *http.Request) {
			var event domain.DeviceEvent

			if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
				http.Error(w, "Invalid JSON data", http.StatusBadRequest)
				return
			}

			deviceIdCtx := r.Context().Value(middlewares.DeviceIDKey)
			deviceId, ok := deviceIdCtx.(int64)
			fmt.Println("Device ID from context:", deviceId)
			if !ok {
				http.Error(w, "Device ID not found in context", http.StatusUnauthorized)
				return
			}

			// Set CreatedAt if you want server timestamp
			event.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
			event.DeviceID = deviceId
			err := events.PublishDeviceDataEvent(event)
			if err != nil {
				http.Error(w, "Failed to enqueue device event", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Device event received and sent to queue",
			})
		})

	})

}
