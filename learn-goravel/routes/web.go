package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/support"

	"learn-goravel/app/facades"
	"learn-goravel/app/http/controllers"
)

func Web() {
	facades.Route().Get("/", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("welcome.tmpl", map[string]any{
			"version": support.Version,
		})
	})

	facades.Route().Static("public", "./public")

	userController := controllers.NewUserController()

	// List all users
	facades.Route().Get("/users", userController.Index)
	facades.Route().Get("/users/create", userController.Create)
	facades.Route().Post("/users/store", userController.Store)
	// Show single user
	facades.Route().Get("/users/{id}", userController.Show)
	facades.Route().Get("/users/{id}/edit", userController.Edit)
	facades.Route().Post("/users/update", userController.Update)
	facades.Route().Post("/users/delete", userController.Delete)

	testController := controllers.NewTestController()
	facades.Route().Get("/users/test", testController.Index)
}
