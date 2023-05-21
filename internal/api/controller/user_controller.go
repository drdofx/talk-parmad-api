package controller

import (
	"github.com/drdofx/talk-parmad/internal/api/models"
	"github.com/drdofx/talk-parmad/internal/api/services"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreateUser(c *gin.Context)
}

type userController struct {
	service services.UserService
}

func NewUserController(service services.UserService) UserController {
	return &userController{service}
}

func (ctr *userController) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Bad request"})
		return
	}

	if err := ctr.service.CreateUser(&user); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": user})
}
