package user

import (
	"gofr.dev/pkg/gofr"

	"github.com/google/uuid"

	"github.com/user-data/errors"
	"github.com/user-data/models"
	"github.com/user-data/service"
	"github.com/user-data/store"
)

type services struct {
	userStore store.User
}

func NewUserService(userStore store.User) service.User {
	return &services{userStore: userStore}
}

func (s services) Create(ctx *gofr.Context, user *models.User) (*models.User, error) {
	err := createValidate(user)
	if err != nil {
		return nil, err
	}

	resp, err := s.userStore.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s services) duplicateCheck(ctx *gofr.Context, user *models.User) error {
	if user.UserName == "" {
		return nil
	}

	f := &models.Filter{UserName: user.UserName}

	users, err := s.GetAll(ctx, f)
	if err != nil {
		return err
	}

	if len(users) > 0 {
		return errors.AlreadyExist{UserName: user.UserName}
	}

	return nil
}

func (s services) Get(ctx *gofr.Context, userID uuid.UUID) (*models.User, error) {
	return s.userStore.Get(ctx, userID)
}

func (s services) GetAll(ctx *gofr.Context, filter *models.Filter) ([]models.User, error) {
	return s.userStore.GetAll(ctx, filter)
}

func (s services) Update(ctx *gofr.Context, user *models.User) (*models.User, error) {
	err := updateValidate(user)
	if err != nil {
		return nil, err
	}

	err = s.duplicateCheck(ctx, user)
	if err != nil {
		return nil, err
	}

	return s.userStore.Update(ctx, user)
}

func (s services) Delete(ctx *gofr.Context, userID uuid.UUID) error {
	_, err := s.Get(ctx, userID)
	if err != nil {
		return err
	}

	return s.userStore.Delete(ctx, userID)
}
