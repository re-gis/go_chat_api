package users

import (
	"chatapp/database"
	"chatapp/utils"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtKey = os.Getenv("JWT_KEY")

var eUser database.User
var newUser database.User

func RegisterUser(c *gin.Context) {

	if err := c.Bind(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newUser.Email == "" || newUser.Name == "" || newUser.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All credentials are required!"})
		return
	}

	// check if user existed
	if err := database.DB.Where("email = ?", newUser.Email).First(&eUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// user not existing

			// upload the profile pic
			filename, ok := c.Get("filePath")
			if !ok {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Filepath not found!"})
				return
			}

			file, ok := c.Get("file")
			if !ok {
				c.JSON(http.StatusBadRequest, gin.H{"error": "File not found!"})
				return
			}

			imageUrl, err := utils.UploadToCloudinary(file.(multipart.File), filename.(string))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			hashedPass, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while hashing the password..."})
				return
			}

			newUser.Password = string(hashedPass)

			newUser.Profile = string(imageUrl)

			if err := database.DB.Create(&newUser).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while registering the user..."})
				return
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"id":    newUser.ID,
				"email": newUser.Email,
			})

			tokenString, err := token.SignedString([]byte(jwtKey))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error..."})
				return
			}

			newUser.Password = ""
			c.JSON(http.StatusOK, gin.H{"user": newUser, "token": tokenString})

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

/* LOGIN */
func LoginUser(c *gin.Context) {
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while binding data..."})
		return
	}

	if newUser.Email == "" || newUser.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All credentials are required!"})
		return
	}

	if err := database.DB.Where("email =?", newUser.Email).First(&eUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password..."})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(eUser.Password), []byte(newUser.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password..."})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    eUser.ID,
		"email": eUser.Email,
	})

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while generating the token..."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully...", "token": tokenString})
}

func DeleteUserAccount(c *gin.Context) {
	userEmail, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Login to continue..."})
		return
	}

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while binding data..."})
		return
	}

	if newUser.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password required to delete the account..."})
		return
	}

	if err := database.DB.Where("email = ?", userEmail).First(&eUser).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(eUser.Password), []byte(newUser.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect password provided!"})
		return
	}

	// delete the user
	if err := database.DB.Delete(&eUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while deleting the user..."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account deleted successfully..."})
}
