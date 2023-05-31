package lib

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RequestHandler struct {
	Gin *gin.Engine
}

func NewRequestHandler() *RequestHandler {
	gin := gin.Default()

	gin.Use(cors.Default())
	return &RequestHandler{gin}
}
