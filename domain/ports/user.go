package ports

import (
	"github.com/atselvan/bkst/domain/models"
	"github.com/google/uuid"
)

type UserService interface {
	FindByEmail(email string) (*models.User, error)
	FindByID(id uuid.UUID) (*models.User, error)
}

type UserRepository interface {
	FindByEmail(email string) (*models.User, error)
	FindByID(id uuid.UUID) (*models.User, error)
}
