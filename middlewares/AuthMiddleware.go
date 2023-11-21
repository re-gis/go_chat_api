package middlewares

import (
	"chatapp/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authotization header not provided..."})
			return
		}

		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) != 2 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid token provided..."})
			return
		}

		tokenString := splitToken[1]
		email, err := utils.ParseToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.Set("email", email)
		c.Next()
	}
}
