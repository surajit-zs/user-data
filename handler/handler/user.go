package handler

import (
	"github.com/google/uuid"

	"gofr.dev/pkg/gofr"

	"github.com/user-data/models"
	"github.com/user-data/service"
)

type Handler struct {
	userService service.User
}

func NewUserHandler(userService service.User) *Handler {
	return &Handler{userService: userService}
}

func (h *Handler) Get(ctx *gofr.Context) (any, error) {
	userID := ctx.PathParam("id")

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	return h.userService.Get(ctx, userUUID)
}

func (h *Handler) Update(ctx *gofr.Context) (any, error) {
	userID := ctx.PathParam("id")

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	var u models.User

	err = ctx.Bind(&u)
	if err != nil {
		return nil, err
	}

	u.ID = userUUID

	resp, err := h.userService.Update(ctx, &u)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *Handler) Create(ctx *gofr.Context) (any, error) {
	var u models.User

	err := ctx.Bind(&u)
	if err != nil {
		return nil, err
	}

	return h.userService.Create(ctx, &u)
}

func (h *Handler) GetAll(ctx *gofr.Context) (any, error) {
	return h.userService.GetAll(ctx, &models.Filter{})
}

func (h *Handler) Delete(ctx *gofr.Context) (any, error) {
	userID := ctx.PathParam("id")

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	return nil, h.userService.Delete(ctx, userUUID)
}
