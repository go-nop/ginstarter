package repositories

import "gorm.io/gorm"

// Repositories is the collection of repositories.
type Repositories struct {
	User UserRepository
}

// NewRepositories creates a new Repositories.
func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User: NewUserRepository(db),
	}
}
