package routes

import (
	"fpt/database"
	"fpt/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		text := ctx.PostForm("text")

		ok := database.UpdateText1(name, text)
		if ok {
			ctx.JSON(http.StatusOK, gin.H{
				"code": global.Success,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": global.Error,
			})
		}
	}
}
