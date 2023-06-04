package routes

import (
	"github.com/drdofx/talk-parmad/internal/api/constants"
	"github.com/drdofx/talk-parmad/internal/api/controller"
	"github.com/drdofx/talk-parmad/internal/api/lib"
	"github.com/drdofx/talk-parmad/internal/api/middleware"
)

type ThreadRoutes interface {
	Route
}

type threadRoutes struct {
	controller controller.ThreadController
	handler    *lib.RequestHandler
}

func NewThreadRoutes(controller controller.ThreadController, handler *lib.RequestHandler) ThreadRoutes {
	return &threadRoutes{controller, handler}
}

func (r *threadRoutes) Setup() {
	auth := r.handler.Gin.Group(constants.API_PATH + "/thread").Use(middleware.AuthorizeJWT())
	{
		auth.POST("/create", r.controller.CreateThread)
		auth.POST("/vote", r.controller.VoteThread)
		auth.PUT("/edit", r.controller.EditThread)
	}
}
