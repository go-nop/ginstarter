package services

import "github.com/go-nop/ginstarter/internal/repositories"

// Services is the collection of services.
type Services struct {
	User UserService
}

// NewServices creates a new Services.
func NewServices(repo *repositories.Repositories) *Services {
	return &Services{
		User: NewUserService(repo.User),
	}
}
