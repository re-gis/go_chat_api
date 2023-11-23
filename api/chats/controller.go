package chats

import (
	"chatapp/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

var eUser database.User
var user database.User

// get my chats
func GetMyChats(c *gin.Context) {
	user_email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Please login to continue..."})
		return
	}

	// get the user
	if err := database.DB.Where("email =?", user_email).First(&eUser).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}

	var chats []database.Chat
	if err := database.DB.Where("senderId = ? OR receiverId =?", eUser.ID, eUser.ID).Find(&chats).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(chats) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No chats found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"chats": chats})
}

func GetChatMessages(c *gin.Context) {
	user_email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Login to continue..."})
		return
	}

	// get user
	if err := database.DB.Where("email = ?", user_email).First(&eUser).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while binding the data..."})
		return
	}

	if user_email == user.Email {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You can't get your messages..."})
		return
	}


	// get the chat of the both users
	var chat database.Chat
	if err := database.DB.Where("senderId = ? AND receiverId = ?", eUser.ID, user.ID).Or("senderId =? AND receiverId = ?", user.ID, eUser.ID).First(&chat).Error; err != nil {
		// create the chat
		chat.ReceiverId = uint(user.ID)
		chat.SenderId = uint(eUser.ID)

		if err := database.DB.Create(&chat).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating the chat space..."})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Empty chat space created successfully...", "chat": chat})
	}

	// get the messages
	var messages []database.Message
	if err := database.DB.Where("chatId =?", chat.ID).Find(&messages).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Messages not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"chat": chat, "messages": messages})
}
