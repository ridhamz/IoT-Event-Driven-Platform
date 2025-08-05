package commands

import (
	"go-cqrs-api/domain"
	"go-cqrs-api/infrastructure"
)

func HandleCreateUser(user domain.User) error {
	db := infrastructure.GetDB()
	query := `
        INSERT INTO users (first_name, last_name, email, password, created_at)
        VALUES (?, ?, ?, ?, ?)
    `
	_, err := db.Exec(query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.CreatedAt,
	)
	return err

	// Publish event
	//return events.PublishUserCreated(id, name)
}
