package handler

import (
	"GopherGate/internal/storage"
	"GopherGate/pkg/api"
	"GopherGate/pkg/repository"
	"GopherGate/pkg/service"
)

// Handler struct contains the API handlers.
type Handler struct {
	Api *api.UsersAPI // UsersAPI provides handlers for user-related operations.
}

// NewUsersHandler creates a new instance of Handler with initialized API handlers.
func NewUsersHandler() *Handler {
	// Get a handle to the database instance
	dbInstance := storage.GetDB()

	// Initialize user repository with the database instance
	userRepository := repository.NewUsersRepository(dbInstance)

	// Initialize user service with the user repository
	userService := service.NewUsersService(userRepository)

	// Initialize API with the user service
	userAPI := api.NewUsersAPI(userService)

	// Create a new handler instance with the initialized API
	return &Handler{
		Api: userAPI,
	}
}
