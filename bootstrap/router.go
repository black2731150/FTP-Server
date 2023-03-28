package bootstrap

import (
	"fpt/global"
	"fpt/routes"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	rootGroup := router.Group("/")
	routes.SetRootGroupRoutes(rootGroup)

	return router
}

func RunServer() {
	router := setupRouter()
	router.Run("0.0.0.0:" + global.Ftpserver.Config.Ftp.Port)
}
