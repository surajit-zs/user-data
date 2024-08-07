package user

import (
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/user-data/errors"
	"github.com/user-data/models"
)

const (
	maxUpdatableFieldCount = 3
)

func createValidate(u *models.User) error {
	u.ID = uuid.New()
	u.CreatedAt = time.Now().UTC()

	u.Name = strings.TrimSpace(u.Name)
	u.UserName = strings.TrimSpace(u.UserName)
	u.Password = strings.TrimSpace(u.Password)

	missingParam := make([]string, 0)

	if u.Name == "" {
		missingParam = append(missingParam, "name")
	}

	if u.UserName == "" {
		missingParam = append(missingParam, "userName")
	}

	if u.Password == "" {
		missingParam = append(missingParam, "password")
	}

	if len(missingParam) > 0 {
		return errors.MissingParam{Params: missingParam}
	}

	return nil
}

func updateValidate(u *models.User) error {
	u.UpdatedAt = time.Now().UTC()

	u.Name = strings.TrimSpace(u.Name)
	u.UserName = strings.TrimSpace(u.UserName)
	u.Password = strings.TrimSpace(u.Password)

	missingParam := make([]string, 0)

	if u.Name == "" {
		missingParam = append(missingParam, "name")
	}

	if u.UserName == "" {
		missingParam = append(missingParam, "userName")
	}

	if u.Password == "" {
		missingParam = append(missingParam, "password")
	}

	if len(missingParam) == maxUpdatableFieldCount {
		return errors.MissingParam{Params: missingParam}
	}

	return nil
}
