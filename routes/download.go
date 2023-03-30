package routes

import (
	"fpt/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DownLAndFile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// ctx.Header("Content-Type", "application/octet-stream")
		// ctx.Header("Content-Disposition", "attachment; filename=test") // 用来指定下载下来的文件名
		// ctx.Header("Content-Transfer-Encoding", "text")
		// ctx.File("./files/upload.go")

		// var content string
		// content = "hello world, 我是一个文件，" + content

		// ctx.Writer.WriteHeader(http.StatusOK)
		// ctx.Header("Content-Disposition", "attachment; filename=hello.txt")
		// ctx.Header("Content-Type", "application/text/plain")
		// ctx.Header("Accept-Length", fmt.Sprintf("%d", len(content)))
		// ctx.Writer.Write([]byte(content))

		filename := ctx.Param("FileName")

		file, err := common.OpenFile("files/" + filename)
		if err != nil {
			ctx.String(http.StatusOK, "Error")
		}
		defer file.Close()

		filestat, err := file.Stat()
		if err != nil {
			ctx.String(http.StatusOK, "Error")
		}

		len := filestat.Size()

		b := make([]byte, len)
		n, err := file.Read(b)
		if err != nil {
			return
		}

		ctx.Header("Content-Type", "application/octet-stream")
		ctx.Writer.Write(b[:n])

	}
}
