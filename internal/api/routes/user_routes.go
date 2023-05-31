package routes

import (
	"github.com/drdofx/talk-parmad/internal/api/constants"
	"github.com/drdofx/talk-parmad/internal/api/controller"
	"github.com/drdofx/talk-parmad/internal/api/lib"
)

type UserRoutes interface {
	Route
}

type userRoutes struct {
	controller controller.UserController
	handler    *lib.RequestHandler
}

func NewUserRoutes(controller controller.UserController, handler *lib.RequestHandler) UserRoutes {
	return &userRoutes{controller, handler}
}

func (r *userRoutes) Setup() {
	auth := r.handler.Gin.Group(constants.API_PATH)
	{
		auth.POST("/login", r.controller.LoginUser)
		auth.POST("/register", r.controller.CreateUser)
	}
}
