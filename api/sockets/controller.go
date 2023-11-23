package sockets

import (
	"chatapp/database"
	"chatapp/utils"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	socketio "github.com/googollee/go-socket.io"
)

var SocketServer *socketio.Server
var user database.User

func InitSocket() error {
	SocketServer = socketio.NewServer(nil)

	return nil
}

func SocketEvents() {
	SocketServer.OnConnect("/", func(conn socketio.Conn) error {
		userId, err := getUserEmailFromContext(conn)
		if err != nil {
			fmt.Errorf("User not found!")
		}
		conn.SetContext(userId)
		conn.Join(userId)
		go pingClientPeriodically(conn)
		fmt.Println("Connected: ", conn.ID())
		return nil
	})

	SocketServer.OnEvent("/", "message", func(conn socketio.Conn, msg string) {
		var input utils.MessageInput

		if err := json.Unmarshal([]byte(msg), &input); err != nil {
			fmt.Println(err)
			return
		}

		if err := input.Validate(); err != nil {
			fmt.Println(err)
			return
		}

		chatId, _ := strconv.Atoi(input.ChatID)
		receiverId, _ := strconv.Atoi(input.ReceiverID)
		senderId, _ := strconv.Atoi(input.SenderID)

		message := database.Message{
			Body:       input.Message,
			ChatId:     uint(chatId),
			ReceiverId: uint(receiverId),
			SenderId:   uint(senderId),
			Isread:     false,
		}
		fmt.Println("Received message:", msg)

		if err := database.DB.Create(&message).Error; err != nil {
			panic("Error occurred while creating the message: " + err.Error())
		}

		userId, err := getUserEmailFromContext(conn)
		if err != nil {
			fmt.Errorf("User not found!")
		}

		event := "message" + input.ReceiverID
		result, _ := json.Marshal(message)
		resultString := string(result)
		conn.SetContext("")
		SocketServer.BroadcastToRoom("/", userId, event, resultString)
		conn.Emit("message"+input.SenderID, resultString)
	})

	SocketServer.OnEvent("/", "bye", func(s socketio.Conn, msg string) {
		fmt.Println(msg)
		log.Println(s.Close())
	})

	SocketServer.OnError("/", func(c socketio.Conn, err error) {
		fmt.Println("Meet the error: ", err)
	})

	SocketServer.OnDisconnect("/", func(c socketio.Conn, reason string) {
		fmt.Println("Closed: ", reason)
	})
}

func getUserEmailFromContext(c socketio.Conn) (string, error) {
	fmt.Println("Context:", c.Context())
	fmt.Println("Conn:", c)
	user_email, exists := c.Context().(string)
	if !exists {
		return "", fmt.Errorf("Please login to continue!")
	}

	// get user id
	if err := database.DB.Where("email =?", user_email).First(&user).Error; err != nil {
		return "", fmt.Errorf("User not found!")
	}

	return strconv.FormatUint(uint64(user.ID), 10), nil
}


func pingClientPeriodically(conn socketio.Conn) {
    ticker := time.NewTicker(10 * time.Second)

    for range ticker.C {
        conn.Emit("ping", )
    }
}