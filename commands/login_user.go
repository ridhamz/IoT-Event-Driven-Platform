package commands

import (
	"go-cqrs-api/domain"
	"go-cqrs-api/infrastructure"
	"go-cqrs-api/utils"

	"golang.org/x/crypto/bcrypt"
)

func HandleLoginUser(req domain.LoginRequest) (string, error) {
	db := infrastructure.GetDB()
	var user domain.User

	err := db.QueryRow(`SELECT id, password FROM users WHERE email = ?`, req.Email).
		Scan(&user.ID, &user.Password)

	if err != nil {
		return "", err
	}

	// Compare password hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", err
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, req.Email)
	if err != nil {
		return "", err
	}

	return token, nil
	// Publish event
	//return events.PublishUserCreated(id, name)
}
