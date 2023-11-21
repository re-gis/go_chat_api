package database

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(c *gin.Context) {
	dsn := "host=localhost user=postgres password=Password@2001 dbname=gin-chatapp port=5432 sslmode=disable"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error..."})
		panic("Failed to connect to the database")
	}

	DB = connection
	DB.AutoMigrate(&User{})
}
