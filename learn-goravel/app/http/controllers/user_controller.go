package controllers

import (
	"learn-goravel/app/facades"
	"learn-goravel/app/models"

	"github.com/goravel/framework/contracts/http"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

// Show create form
func (r *UserController) Create(ctx http.Context) http.Response {
	data := map[string]any{
		"PageTitle": "Create New User",
	}
	return ctx.Response().View().Make("create.tmpl", data)
}

// Store new user
func (r *UserController) Store(ctx http.Context) http.Response {
	var user models.Users

	user.Name = ctx.Request().Input("name")
	user.Email = ctx.Request().Input("email")
	user.Password = ctx.Request().Input("password")

	age := ctx.Request().InputInt("age")
	user.Age = age

	city := ctx.Request().Input("city")
	user.City = city

	isActive := ctx.Request().InputBool("is_active")
	user.IsActive = isActive

	facades.Orm().Query().Create(&user)

	return ctx.Response().Json(http.StatusOK, http.Json{
		"message": "User created successfully",
		"user":    user,
	})
}

// Show all users
func (r *UserController) Index(ctx http.Context) http.Response {
	var users []models.Users

	facades.Orm().Query().Find(&users)

	data := map[string]any{
		"PageTitle": "All Users",
		"Users":     users,
		"Count":     len(users),
	}

	return ctx.Response().View().Make("index.tmpl", data)
}

// Show single user
func (r *UserController) Show(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")

	var user models.Users
	facades.Orm().Query().Find(&user, id)

	data := map[string]any{
		"PageTitle": "User Profile",
		"User":      user,
	}

	return ctx.Response().View().Make("show.tmpl", data)
}

func (r *UserController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Input("id")

	var user models.Users
	facades.Orm().Query().Find(&user, id)

	user.Name = ctx.Request().Input("name")
	user.Email = ctx.Request().Input("email")

	age := ctx.Request().InputInt("age")
	user.Age = age

	city := ctx.Request().Input("city")
	user.City = city

	isActive := ctx.Request().InputBool("is_active")
	user.IsActive = isActive

	facades.Orm().Query().Save(&user)

	return ctx.Response().Json(http.StatusOK, http.Json{
		"message": "User updated successfully",
		"user":    user,
	})
}

func (r *UserController) Edit(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")

	var user models.Users
	facades.Orm().Query().Find(&user, id)

	data := map[string]any{
		"PageTitle": "Edit User",
		"User":      user,
	}
	return ctx.Response().View().Make("edit.tmpl", data)
}

func (r *UserController) Delete(ctx http.Context) http.Response {
	id := ctx.Request().Input("id")

	var user models.Users
	facades.Orm().Query().Find(&user, id)

	facades.Orm().Query().Delete(&user)

	return ctx.Response().Json(http.StatusOK, http.Json{
		"message": "User deleted successfully",
	})
}
