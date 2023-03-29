package routes

import (
	"fpt/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		text := ctx.PostForm("text")

		ok := database.UpdateText(name, text)
		if ok {
			ctx.JSON(http.StatusOK, gin.H{
				"code": "SUCCESS",
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": "ERROR",
			})
		}
	}
}
