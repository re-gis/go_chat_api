package routes

import (
	"chatapp/api/sockets"
	"chatapp/api/users"
	"chatapp/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) *gin.Engine {

	r.POST("/users/register", middlewares.FileUploadMiddleware(), users.RegisterUser)
	r.POST("/users/login", users.LoginUser)

	protected := r.Group("/")
	protected.Use(middlewares.Authentication())

	SetupUserRoutes(protected)
	SetupSocketRoutes(protected)

	return r
}

func SetupUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/users")
	{
		userRoute.DELETE("/delete-account", users.DeleteUserAccount)
		userRoute.PATCH("/update/:id", users.UpdateUserAccount)
	}
}

func SetupSocketRoutes(rg *gin.RouterGroup) {
	socketRouter := rg.Group("/socket.io")

	socketRouter.GET("/*any", gin.WrapH(sockets.SocketServer))
	socketRouter.POST("/*any", gin.WrapH(sockets.SocketServer))
}

