package main

import (
	"gofr.dev/pkg/gofr"

	"github.com/user-data/handler/handler"
	"github.com/user-data/migrations"
	service "github.com/user-data/service/user"
	store "github.com/user-data/store/user"
)

func main() {
	app := gofr.New()

	app.Migrate(migrations.All())

	//app.EnableAPIKeyAuthWithValidator()

	userStore := store.NewUserStore()
	userService := service.NewUserService(userStore)
	userHandler := handler.NewUserHandler(userService)

	app.GET("/v1/user-data", userHandler.GetAll)
	app.GET("/v1/user-data/{id}", userHandler.Get)
	app.POST("/v1/user-data", userHandler.Create)
	app.PATCH("/v1/user-data/{id}", userHandler.Update)
	app.DELETE("/v1/user-data/{id}", userHandler.Delete)

	app.Run()
}
