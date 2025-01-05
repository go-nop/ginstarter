package services

import (
	"context"

	"github.com/go-nop/ginstarter/internal/entity"
	"github.com/go-nop/ginstarter/internal/repositories"
)

// UserService is the service for users.
type UserService interface {
	FindUserByID(ctx context.Context, id string) (*entity.User, error)
	SaveUser(ctx context.Context, user *entity.User) error
	GetUsers(ctx context.Context) ([]entity.User, error)
}

// userService is the implementation of UserService.
type userService struct {
	userRepository repositories.UserRepository
}

// UserServiceInterface is the interface for UserService.
var _ UserService = &userService{}

// NewUserService creates a new UserService.
func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

// FindUserByID finds a user by ID.
func (s *userService) FindUserByID(ctx context.Context, id string) (*entity.User, error) {
	user, err := s.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// SaveUser saves a user.
func (s *userService) SaveUser(ctx context.Context, user *entity.User) error {
	if err := s.userRepository.Save(ctx, user); err != nil {
		return err
	}

	return nil
}

// GetUsers gets all users.
func (s *userService) GetUsers(ctx context.Context) ([]entity.User, error) {
	users, err := s.userRepository.Get(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
