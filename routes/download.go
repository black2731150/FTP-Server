package routes

import (
	"fpt/common"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DownoadFile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		filename := ctx.Param("FileName")

		file, err := common.OpenFile("files/" + filename)
		if err != nil {
			ctx.String(http.StatusOK, "Error")
		}
		defer file.Close()

		chuck := 1024 * 4

		ctx.Header("Content-Type", "application/octet-stream")

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
