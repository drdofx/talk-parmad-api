package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleSuccessResponse(c *gin.Context, data interface{}) {
	res := Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    data,
	}
	c.JSON(http.StatusOK, res)
}

func HandleErrorResponse(c *gin.Context, status int, message string) {
	res := Response{
		Status:  status,
		Message: message,
		Data:    nil,
	}
	c.JSON(status, res)
}
