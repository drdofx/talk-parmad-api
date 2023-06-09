package routes

import (
	"github.com/drdofx/talk-parmad/internal/api/constants"
	"github.com/drdofx/talk-parmad/internal/api/controller"
	"github.com/drdofx/talk-parmad/internal/api/lib"
	"github.com/drdofx/talk-parmad/internal/api/middleware"
)

type ForumRoutes interface {
	Route
}

type forumRoutes struct {
	controller controller.ForumController
	handler    *lib.RequestHandler
}

func NewForumRoutes(controller controller.ForumController, handler *lib.RequestHandler) ForumRoutes {
	return &forumRoutes{controller, handler}
}

func (r *forumRoutes) Setup() {
	auth := r.handler.Gin.Group(constants.API_PATH + "/forum").Use(middleware.AuthorizeJWT())
	{
		auth.POST("/create", r.controller.CreateForum)
		auth.POST("/join", r.controller.JoinForum)
		auth.PUT("/edit", r.controller.EditForum)
		auth.DELETE("/delete", r.controller.DeleteForum)
		auth.GET("/list", r.controller.ListUserForum)
		auth.GET("/discover", r.controller.ListDiscoverForum)
		auth.GET("/detail", r.controller.DetailForum)
		auth.GET("/list-thread", r.controller.ListThreadForumHome)
		auth.PUT("/remove", r.controller.RemoveFromForum)
		auth.GET("/search", r.controller.SearchForum)
	}
}
