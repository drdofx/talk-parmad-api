package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/drdofx/talk-parmad/internal/api/helper"
	"github.com/drdofx/talk-parmad/internal/api/lib"
	"github.com/drdofx/talk-parmad/internal/api/request"
	"github.com/drdofx/talk-parmad/internal/api/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ThreadController interface {
	CreateThread(c *gin.Context)
	VoteThread(c *gin.Context)
	EditThread(c *gin.Context)
	DetailThread(c *gin.Context)
	CreateReply(c *gin.Context)
	VoteReply(c *gin.Context)
	EditReply(c *gin.Context)
	DeleteThread(c *gin.Context) // only moderator
	DeleteReply(c *gin.Context)  // only moderator
}

type threadController struct {
	services services.ThreadService
	validate *validator.Validate
}

func NewThreadController(service services.ThreadService, validate *validator.Validate) ThreadController {
	return &threadController{service, validate}
}

func (ctr *threadController) CreateThread(c *gin.Context) {
	var req request.ReqSaveThread

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

	res, err := ctr.services.CreateThread(&req, &user)

	if err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.HandleSuccessResponse(c, res)
}

func (ctr *threadController) VoteThread(c *gin.Context) {
	var req request.ReqVoteThread

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

	fmt.Println("req.Vote: ", req.Vote)
	user := helper.GetUserData(c)

	res, err := ctr.services.VoteThread(&req, &user)

	if err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.HandleSuccessResponse(c, res)
}

func (ctr *threadController) EditThread(c *gin.Context) {
	var req request.ReqEditThread

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

	res, err := ctr.services.EditThread(&req, &user)

	if err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.HandleSuccessResponse(c, res)
}

func (ctr *threadController) DetailThread(c *gin.Context) {
	var req request.ReqDetailThread

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

	res, err := ctr.services.DetailThread(&req)

	if err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.HandleSuccessResponse(c, res)
}

// Reply controller
func (ctr *threadController) CreateReply(c *gin.Context) {
	var req request.ReqSaveReply

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

	res, err := ctr.services.CreateReply(&req, &user)

	if err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.HandleSuccessResponse(c, res)
}

func (ctr *threadController) VoteReply(c *gin.Context) {
	var req request.ReqVoteReply

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

	res, err := ctr.services.VoteReply(&req, &user)

	if err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.HandleSuccessResponse(c, res)
}

func (ctr *threadController) EditReply(c *gin.Context) {
	var req request.ReqEditReply

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

	res, err := ctr.services.EditReply(&req, &user)

	if err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.HandleSuccessResponse(c, res)
}

// MODERATOR ONLY CONTROLLERS
func (ctr *threadController) DeleteThread(c *gin.Context) {
	var req request.ReqDeleteThread

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

	threadIdInt, _ := strconv.Atoi(req.ThreadID)
	thread, err := ctr.services.GetThreadByID(uint(threadIdInt))
	if err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	_, err = ctr.services.CheckModeratorForumFromThread(thread, &user)
	if err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = ctr.services.DeleteThread(thread)

	if err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.HandleSuccessResponse(c, nil)
}

func (ctr *threadController) DeleteReply(c *gin.Context) {
	var req request.ReqDeleteReply

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

	replyIdInt, _ := strconv.Atoi(req.ReplyID)
	thread, reply, err := ctr.services.GetThreadAndReplyByReplyID(uint(replyIdInt))
	if err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	_, err = ctr.services.CheckModeratorForumFromThread(thread, &user)
	if err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = ctr.services.DeleteReply(reply)

	if err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.HandleSuccessResponse(c, nil)
}
