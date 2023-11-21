package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FileUploadMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		file, header, err := ctx.Request.FormFile("file")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad request..."})
			return
		}

		defer file.Close()
		ctx.Set("filePath", header.Filename)
		ctx.Set("file", file)
		ctx.Next()
	}
}
