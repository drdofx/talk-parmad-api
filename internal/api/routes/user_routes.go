package routes

import (
	"github.com/drdofx/talk-parmad/internal/api/controller"
	"github.com/gin-gonic/gin"
)

type UserRoutes interface {
	SetupUserRoutes(router *gin.RouterGroup)
}

type userRoutes struct {
	controller controller.UserController
}

func NewUserRoutes(controller controller.UserController) UserRoutes {
	return &userRoutes{controller}
}

func (r *userRoutes) SetupUserRoutes(router *gin.RouterGroup) {
	router.POST("/users", r.controller.CreateUser)
}
