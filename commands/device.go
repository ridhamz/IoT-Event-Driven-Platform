package commands

import (
	"go-cqrs-api/domain"
	"go-cqrs-api/infrastructure"
	"go-cqrs-api/utils"
	"log"
)

func HandleCreateDevice(device domain.Device) error {
	db := infrastructure.GetDB()

	query := `
		INSERT INTO devices (name, user_id, created_at)
		VALUES (?, ?, ?)
	`

	result, err := db.Exec(query, device.Name, device.UserID, device.CreatedAt)
	if err != nil {
		log.Println("Error inserting device:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err == nil {
		device.ID = id
	}

	utils.GenerateAndStoreAPIKey(device.ID)

	return nil
}
