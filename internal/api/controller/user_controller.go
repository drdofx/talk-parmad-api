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

type UserController interface {
	CreateUser(c *gin.Context)
	LoginUser(c *gin.Context)
}

type userController struct {
	service  services.UserService
	validate *validator.Validate
}

func NewUserController(service services.UserService, validate *validator.Validate) UserController {
	return &userController{service, validate}
}

func (ctr *userController) CreateUser(c *gin.Context) {
	var req request.ReqSaveUser

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

	res, err := ctr.service.CreateUser(&req)

	if err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.HandleSuccessResponse(c, res)
}

func (ctr *userController) LoginUser(c *gin.Context) {
	var req request.ReqLoginUser

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

	res, err := ctr.service.LoginUser(&req)

	if err != nil {
		lib.CommonLogger().Error(err)
		helper.HandleErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.HandleSuccessResponse(c, res)
}
