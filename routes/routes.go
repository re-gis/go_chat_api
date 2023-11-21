package routes

import (
	"chatapp/api/users"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) *gin.Engine {
	r.POST("/users/register", users.RegisterUser)

	return r
}
