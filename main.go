package main

import (
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


	// database
	database.Connect()

	/* ROUTES */
	routes.SetupRoutes(r)

	if port == "" {
		port = "8080"
	}

	r.Run()
}
