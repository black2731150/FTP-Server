package bootstrap

import (
	"fpt/global"
	"fpt/routes"

	"github.com/gin-gonic/gin"
)

//路由管理
func setupRouter() *gin.Engine {
	router := gin.Default()

	rootGroup := router.Group("/")
	routes.SetRootGroupRoutes(rootGroup)

	return router
}

//启动服务
func RunServer() {
	router := setupRouter()
	router.Run(global.Ftpserver.Config.Ftp.Host + ":" + global.Ftpserver.Config.Ftp.Port)
}
