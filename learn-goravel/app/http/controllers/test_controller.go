package controllers

import (
	"github.com/goravel/framework/contracts/http"
)

type TestController struct{}

func NewTestController() *TestController {
	return &TestController{}
}

func (r *TestController) Index(ctx http.Context) http.Response {
	return ctx.Response().Success().Json(http.Json{
		"Test": "Finally Goravel Working Started",
	})
}
