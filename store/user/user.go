package user

import (
	"database/sql"
	"time"

	"github.com/google/uuid"

	"gofr.dev/pkg/gofr"

	errorsPkg "github.com/pkg/errors"

	"github.com/user-data/errors"
	"github.com/user-data/models"
	"github.com/user-data/store"
)

type Store struct{}

func NewUserStore() store.User {
	return Store{}
}

func (Store) Create(ctx *gofr.Context, user *models.User) (*models.User, error) {
	_, err := ctx.SQL.ExecContext(ctx, createUserQuery, user.ID, user.Name, user.UserName,
		user.Password, user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (Store) Get(ctx *gofr.Context, userID uuid.UUID) (*models.User, error) {
	var (
		u        models.User
		updateAt sql.NullTime
	)

	err := ctx.SQL.QueryRowContext(ctx, getUserQuery, userID).Scan(&u.ID,
		&u.Name, &u.UserName, &u.Password, &u.CreatedAt, &updateAt)
	if err != nil {
		if errorsPkg.Is(err, sql.ErrNoRows) {
			return nil, errors.NotFound{ID: userID.String()}
		}

		return nil, errors.DB{Err: err}
	}

	u.UpdatedAt = updateAt.Time

	return &u, err
}

func (Store) GetAll(ctx *gofr.Context, filter *models.Filter) ([]models.User, error) {
	query, args := getAllBuildQuery(filter)

	rows, err := ctx.SQL.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	defer rows.Close()

	users := make([]models.User, 0)

	for rows.Next() {
		var (
			u        models.User
			updateAt sql.NullTime
		)

		err = rows.Scan(&u.ID, &u.Name, &u.UserName, &u.Password, &u.CreatedAt, &updateAt)
		if err != nil {
			return nil, errors.DB{Err: err}
		}

		u.UpdatedAt = updateAt.Time

		users = append(users, u)
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	return users, nil
}

func (s Store) Update(ctx *gofr.Context, user *models.User) (*models.User, error) {
	query, args := buildUpdateQuery(user)

	_, err := ctx.SQL.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	return s.Get(ctx, user.ID)
}

func (Store) Delete(ctx *gofr.Context, userID uuid.UUID) error {
	_, err := ctx.SQL.ExecContext(ctx, deleteUserQuery, time.Now().UTC(), userID)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}
