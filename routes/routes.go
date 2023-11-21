package routes

import (
	"chatapp/api/users"
	"chatapp/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) *gin.Engine {

	r.POST("/users/register", middlewares.FileUploadMiddleware(), users.RegisterUser)
	r.POST("/users/login", users.LoginUser)
	r.DELETE("/users/delete-account", middlewares.Authentication(), users.DeleteUserAccount)

	return r
}
