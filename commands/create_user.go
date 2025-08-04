package commands

import (
	"go-cqrs-api/events"
	"go-cqrs-api/infrastructure"
)

func HandleCreateUser(id, name string) error {
	db := infrastructure.GetDB()
	_, err := db.Exec(`INSERT INTO users (id, name) VALUES (?, ?)`, id, name)
	if err != nil {
		return err
	}

	// Publish event
	return events.PublishUserCreated(id, name)
}
