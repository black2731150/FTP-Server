package main

import (
	"fpt/bootstrap"
	"fpt/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	bootstrap.InitializeConfig()

	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})

	router.Run("0.0.0.0" + ":" + global.Ftp.Config.Ftp.Port)
}
