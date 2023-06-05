package controller

import (
	"fmt"
	"net/http"

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
	// GetThreads(c *gin.Context)
	// UpdateThread(c *gin.Context)
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