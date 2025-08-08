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

func GetUserDevices(userID int64) ([]domain.DeviceWithAPIKeys, error) {
	db := infrastructure.GetDB()

	// Query devices
	devicesQuery := `
		SELECT id, name, user_id, created_at
		FROM devices
		WHERE user_id = ?
	`
	rows, err := db.Query(devicesQuery, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var devices []domain.DeviceWithAPIKeys

	for rows.Next() {
		var d domain.DeviceWithAPIKeys
		if err := rows.Scan(&d.ID, &d.Name, &d.UserID, &d.CreatedAt); err != nil {
			return nil, err
		}

		// Query API keys for this device
		keysQuery := `
			SELECT id, api_key, created_at
			FROM device_api_keys
			WHERE device_id = ?
		`
		keyRows, err := db.Query(keysQuery, d.ID)
		if err != nil {
			return nil, err
		}

		var apiKeys []domain.APIKey
		for keyRows.Next() {
			var k domain.APIKey
			if err := keyRows.Scan(&k.ID, &k.APIKey, &k.CreatedAt); err != nil {
				keyRows.Close()
				return nil, err
			}
			apiKeys = append(apiKeys, k)
		}
		keyRows.Close()

		d.APIKeys = apiKeys
		devices = append(devices, d)
	}

	return devices, nil
}
