package router

import (
	"IMChat_App/internal/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter() {
	r := gin.Default()
	r.POST("/user/register", handler.Register)
	r.Run(":9090")
}
