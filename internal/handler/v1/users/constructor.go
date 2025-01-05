package users

import "github.com/go-nop/ginstarter/internal/services"

// UserHandler is the handler for users.
type UserHandler struct {
	UserService services.UserService
}

// NewUserHandler creates a new UserHandler.
func NewHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}
