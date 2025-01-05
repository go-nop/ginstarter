package repositories

import (
	"context"

	"github.com/go-nop/ginstarter/internal/entity"
	"gorm.io/gorm"
)

// UserRepository is the repository for users.
type UserRepository interface {
	// FindByID finds a user by ID.
	FindByID(ctx context.Context, id string) (*entity.User, error)
	// Save saves a user.
	Save(ctx context.Context, user *entity.User) error
	// Get gets all users.
	Get(ctx context.Context) ([]entity.User, error)
	// DeleteByID deletes a user.
	DeleteByID(ctx context.Context, id string) error
}

// userRepository is the implementation of UserRepository.
type userRepository struct {
	db *gorm.DB
}

// UserRepositoryInterface is the interface for UserRepository.
var _ UserRepository = &userRepository{}

// NewUserRepository creates a new UserRepository.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// FindByID finds a user by ID.
func (r *userRepository) FindByID(ctx context.Context, id string) (*entity.User, error) {
	var user entity.User

	if err := r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", id).Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Save saves a user.
func (r *userRepository) Save(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).Save(user).Error; err != nil {
		return err
	}

	return nil
}

// Get gets all users.
func (r *userRepository) Get(ctx context.Context) ([]entity.User, error) {
	var users []entity.User

	if err := r.db.WithContext(ctx).Model(&entity.User{}).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// DeleteByID soft deletes a user by ID.
func (r *userRepository) DeleteByID(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Delete(&entity.User{}, id).Error; err != nil {
		return err
	}

	return nil
}
