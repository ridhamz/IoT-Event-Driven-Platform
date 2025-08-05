package api

import (
	"encoding/json"
	"fmt"
	"go-cqrs-api/commands"
	"go-cqrs-api/domain"
	"go-cqrs-api/queries"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
)

func SetupRouter() http.Handler {
	r := chi.NewRouter()

	r.Post("/users", func(w http.ResponseWriter, r *http.Request) {
		var body domain.User
		fmt.Print(r.Body)
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
			return
		}

		user := domain.User{
			ID:        body.ID,
			FirstName: body.FirstName,
			LastName:  body.LastName,
			Email:     body.Email,
			Password:  string(hashedPassword),
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		fmt.Println("User to be created:", user)
		if err := commands.HandleCreateUser(user); err != nil {
			fmt.Println("Error creating user:", err)
			http.Error(w, "Could not create user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})

	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		var req domain.LoginRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		token, err := commands.HandleLoginUser(req)
		if err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		// Return token in response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"token": token,
		})

	})

	r.Get("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		user, err := queries.GetUserFromReadModel(id)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(user)
	})

	return r
}
