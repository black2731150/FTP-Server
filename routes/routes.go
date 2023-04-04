package routes

import (
	"fmt"
	"fpt/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

//网站根目录下的路由集合
func SetRootGroupRoutes(router *gin.RouterGroup) {
	//静态文件服务
	router.StaticFS("/web", http.Dir("./web"))
	router.StaticFile("/", "./web/index.html")

	//下载文件服务
	router.GET(fmt.Sprintf("/%s/:FileName", global.Ftpserver.Config.Ftp.Storage), DownoadFile())

	//上传文件服务
	router.POST("/upload", Upload())

	//前端获取文件信息接口
	router.GET("/getFileInfo", GetFileInfo())

	//备注更新接口
	router.POST("/update", Update())
}
