package controller

import (
	"fmt"

	"github.com/drdofx/talk-parmad/internal/api/helper"
	"github.com/drdofx/talk-parmad/internal/api/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ForumController interface {
	CreateForum(c *gin.Context)
}

type forumController struct {
	services services.ForumService
	validate *validator.Validate
}

func NewForumController(service services.ForumService, validate *validator.Validate) ForumController {
	return &forumController{service, validate}
}

func (ctr *forumController) CreateForum(c *gin.Context) {
	user := helper.GetUserData(c)

	fmt.Println(user)

	helper.HandleSuccessResponse(c, "OK")
}
