package routes

import (
	"github.com/drdofx/talk-parmad/internal/api/controller"
	"github.com/gin-gonic/gin"
)

type UserRoutes interface {
	SetupAuthRoutes(router *gin.RouterGroup)
}

type userRoutes struct {
	controller controller.UserController
}

func NewUserRoutes(controller controller.UserController) UserRoutes {
	return &userRoutes{controller}
}

func (r *userRoutes) SetupAuthRoutes(router *gin.RouterGroup) {
	router.POST("/login", r.controller.LoginUser)
	router.POST("/register", r.controller.CreateUser)
}
