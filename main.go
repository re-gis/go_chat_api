package main

import (
	"chatapp/api/sockets"
	"chatapp/database"
	"chatapp/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading .env file...")
	}

	port := os.Getenv("PORT")

	r := gin.Default()

	if err := sockets.InitSocket(); err != nil {
		panic(err)
	}

	// database
	database.Connect()

	/* ROUTES */
	routes.SetupRoutes(r)

	sockets.SocketEvents()
	go sockets.SocketServer.Serve()
	defer sockets.SocketServer.Close()

	if port == "" {
		port = "8080"
	}

	r.Run()
}
