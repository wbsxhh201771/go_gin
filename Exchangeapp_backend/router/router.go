package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/api/auth")
	{
		auth.POST("/login", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK,gin.H{
				"msg":"login sucess",
			})
		})
		auth.POST("/register",func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK,gin.H{
				"msg":"register sucess",
			})
		})
	}
	return r
}