package services

import (
	"github.com/atselvan/bkst/domain/models"
	"github.com/atselvan/bkst/domain/ports"
	"github.com/google/uuid"
)

type userService struct {
	repo ports.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(repo ports.UserRepository) ports.UserService {
	return &userService{
		repo: repo,
	}
}

func (u userService) FindByEmail(email string) (*models.User, error) {
	return u.repo.FindByEmail(email)
}

func (u userService) FindByID(id uuid.UUID) (*models.User, error) {
	return u.repo.FindByID(id)
}
