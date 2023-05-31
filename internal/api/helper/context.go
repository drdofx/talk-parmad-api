package helper

import (
	"github.com/drdofx/talk-parmad/internal/api/lib"
	"github.com/gin-gonic/gin"
)

func GetUserData(c *gin.Context) *lib.UserData {
	user := c.MustGet("USER_DATA")
	return user.(*lib.UserData)
}
