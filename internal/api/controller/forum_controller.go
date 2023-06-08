package controller

import (
	"net/http"

	"github.com/drdofx/talk-parmad/internal/api/helper"
	"github.com/drdofx/talk-parmad/internal/api/lib"
	"github.com/drdofx/talk-parmad/internal/api/request"
	"github.com/drdofx/talk-parmad/internal/api/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ForumController interface {
	CreateForum(c *gin.Context)
	JoinForum(c *gin.Context)
	ListUserForum(c *gin.Context)
	DetailForum(c *gin.Context)
	ListThreadForumHome(c *gin.Context)
	EditForum(c *gin.Context)       // only moderator
	DeleteForum(c *gin.Context)     // only moderator
	RemoveFromForum(c *gin.Context) // only moderator
}

type forumController struct {
	services services.ForumService
	validate *validator.Validate
}

func NewForumController(service services.ForumService, validate *validator.Validate) ForumController {
	return &forumController{service, validate}
}

func (ctr *forumController) CreateForum(c *gin.Context) {
	var req request.ReqSaveForum

	if err := c.ShouldBindJSON(&req); err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, "Bad request")
		return
	}

	if err := ctr.validate.Struct(&req); err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, "Bad input")
		return
	}

	user := helper.GetUserData(c)

	res, err := ctr.services.CreateForum(&req, &user)

	if err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.HandleSuccessResponse(c, res)
}

func (ctr *forumController) JoinForum(c *gin.Context) {
	var req request.ReqJoinForum

	if err := c.ShouldBindJSON(&req); err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, "Bad request")
		return
	}

	if err := ctr.validate.Struct(&req); err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, "Bad input")
		return
	}

	user := helper.GetUserData(c)

	err := ctr.services.JoinForum(&req, &user)

	if err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.HandleSuccessResponse(c, nil)
}

func (ctr *forumController) ListUserForum(c *gin.Context) {
	user := helper.GetUserData(c)

	res, err := ctr.services.ListUserForum(&user)

	if err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.HandleSuccessResponse(c, res)
}

func (ctr *forumController) DetailForum(c *gin.Context) {
	var req request.ReqDetailForum

	if err := c.ShouldBindJSON(&req); err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, "Bad request")
		return
	}

	if err := ctr.validate.Struct(&req); err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, "Bad input")
		return
	}

	// user := helper.GetUserData(c)

	res, err := ctr.services.DetailForum(&req)

	if err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.HandleSuccessResponse(c, res)
}

func (ctr *forumController) ListThreadForumHome(c *gin.Context) {
	user := helper.GetUserData(c)

	res, err := ctr.services.ListThreadForumHome(&user)

	if err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.HandleSuccessResponse(c, res)
}

// MODERATOR ONLY CONTROLLERS
func (ctr *forumController) EditForum(c *gin.Context) {
	var req request.ReqEditForum

	if err := c.ShouldBindJSON(&req); err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, "Bad request")
		return
	}

	if err := ctr.validate.Struct(&req); err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, "Bad input")
		return
	}

	user := helper.GetUserData(c)

	_, err := ctr.services.CheckModeratorForum(&request.ReqCheckModeratorForum{ForumID: req.ForumID, UserID: user.UserID})

	if err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := ctr.services.EditForum(&req)

	if err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.HandleSuccessResponse(c, res)
}

func (ctr *forumController) DeleteForum(c *gin.Context) {
	var req request.ReqDeleteForum

	if err := c.ShouldBindJSON(&req); err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, "Bad request")
		return
	}

	if err := ctr.validate.Struct(&req); err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, "Bad input")
		return
	}

	user := helper.GetUserData(c)

	_, err := ctr.services.CheckModeratorForum(&request.ReqCheckModeratorForum{ForumID: req.ForumID, UserID: user.UserID})

	if err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = ctr.services.DeleteForum(&req)

	if err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.HandleSuccessResponse(c, nil)
}

func (ctr *forumController) RemoveFromForum(c *gin.Context) {
	var req request.ReqRemoveFromForum

	if err := c.ShouldBindJSON(&req); err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, "Bad request")
		return
	}

	if err := ctr.validate.Struct(&req); err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, "Bad input")
		return
	}

	user := helper.GetUserData(c)

	_, err := ctr.services.CheckModeratorForum(&request.ReqCheckModeratorForum{ForumID: req.ForumID, UserID: user.UserID})

	if err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = ctr.services.RemoveFromForum(&req)

	if err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.HandleSuccessResponse(c, nil)
}
