package service

import (
	"gofr.dev/pkg/gofr"

	"github.com/google/uuid"

	"github.com/user-data/models"
)

type User interface {
	Create(ctx *gofr.Context, user *models.User) (*models.User, error)
	Get(ctx *gofr.Context, userID uuid.UUID) (*models.User, error)
	GetAll(ctx *gofr.Context, filter *models.Filter) ([]models.User, error)
	Update(ctx *gofr.Context, user *models.User) (*models.User, error)
	Delete(ctx *gofr.Context, userID uuid.UUID) error
}
