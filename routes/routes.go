package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetRootGroupRoutes(router *gin.RouterGroup) {
	//测试路由
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

}
