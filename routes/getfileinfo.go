package routes

import (
	"fmt"
	"fpt/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetFileInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
		if err != nil {
			fmt.Println("Get page error: ", err)
		}
		size, err := strconv.Atoi(ctx.DefaultQuery("size", "15"))
		if err != nil {
			fmt.Println("Get size error: ", err)
		}

		search := ctx.DefaultQuery("search", "")

		Size, total, fileinfos := database.GetFileInfoFromDB(page, size, search)

		ctx.JSON(http.StatusOK, gin.H{
			"code":  "SUCCESS", //表示请求过去数据成功
			"data":  fileinfos, //请求的数据
			"total": total,     //总数据的数量
			"size":  Size,      //搜索返回的数据数量
		})

	}
}
