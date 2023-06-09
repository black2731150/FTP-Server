package routes

import (
	"fmt"
	"fpt/database"
	"fpt/global"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Upload() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, fmt.Sprintf("%s\n%s", ctx.ContentType(), ctx.Request.RemoteAddr))
		file, err := ctx.FormFile("file")
		// text := ctx.GetHeader("Text")
		// branch := ctx.GetHeader("Branch")
		text := ctx.Query("text")
		branch := ctx.Query("branch")

		if err != nil {
			ctx.String(http.StatusInternalServerError, "error: Get file name error.\n")
		} else {
			pwd, _ := os.Getwd()
			filepath := fmt.Sprintf("%s/%s/%s", pwd, global.Ftpserver.Config.Ftp.Storage, file.Filename)
			ctx.SaveUploadedFile(file, filepath)

			//放入数据库
			err := database.WriteToDB(filepath, text, branch)
			if err != nil {
				fmt.Println("Wirte to DB error: ", err)
			}
		}
	}
}
