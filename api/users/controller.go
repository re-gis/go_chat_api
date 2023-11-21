package users

import (
	"chatapp/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterUser(c *gin.Context) {
	var newUser database.User
	var eUser database.User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newUser.Email == "" || newUser.Name == "" || newUser.Password == "" || newUser.Profile == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All credentials are required!"})
		return
	}

	// check if user existed
	if err := database.DB.Where("email = ?", newUser.Email).First(&eUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// user not existing

			// upload the profile pic
			


			hashedPass, err:=bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
			if err!=nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error":"Error while hashing the password..."})
				return
			}

			newUser.Password = string(hashedPass)

		} else {
			// other error
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error..."})
			return
		}
	} else {
		// user exists
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already taken..."})
		return
	}
}
