package user

import (
	"time"

	"github.com/atselvan/bkst/domain/models"
	"github.com/atselvan/bkst/domain/ports"
	"github.com/google/uuid"
)

type userRepository struct {
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository() ports.UserRepository {
	return &userRepository{}
}

// FindByEmail retrieves a user by their email address
func (u userRepository) FindByEmail(email string) (*models.User, error) {
	return &models.User{
		ID:           uuid.New(),
		Name:         "Allan Tony Selvan",
		Email:        "allantony2008@gmail.com",
		PasswordHash: "password",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, nil
}

func (u userRepository) FindByID(id uuid.UUID) (*models.User, error) {
	return &models.User{
		ID:           uuid.New(),
		Name:         "Allan Tony Selvan",
		Email:        "allantony2008@gmail.com",
		PasswordHash: "password",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, nil
}
