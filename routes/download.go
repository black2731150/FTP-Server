package routes

import (
	"fpt/common"
	"fpt/global"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DownoadFile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		filename := ctx.Param("FileName")
		filepath := global.FileServerFile + filename

		file, err := common.OpenFile(filepath)
		if err != nil {
			ctx.String(http.StatusOK, "Error")
		}
		defer file.Close()

		chuck := 1024 * 4

		ctx.Header(global.ConnectType, global.ConnectTypeDlwnload)

		for {
			buf := make([]byte, chuck)
			n, err := file.Read(buf)
			if err != nil {
				ctx.String(http.StatusOK, "Error")
			}
			if err == io.EOF {
				break
			}
			ctx.Writer.Write(buf[:n])
		}
	}
}
