package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//网站根目录下的路由集合
func SetRootGroupRoutes(router *gin.RouterGroup) {
	//测试路由
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	//静态文件服务
	router.StaticFS("/files", http.Dir("./files"))
	router.StaticFS("/web", http.Dir("./web"))
	router.StaticFile("/", "./web/index.html")

	//上传文件服务
	router.POST("./upload", Upload())

}
