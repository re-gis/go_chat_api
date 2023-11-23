package utils

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type MessageInput struct {
	SenderID   string `json:"senderId"`
	ReceiverID string `json:"receiverId"`
	Message    string `json:"message"`
	ChatID     string `json:"chatId"`
}

// Validate messages
func (msgInput MessageInput) Validate() error {
	return validation.ValidateStruct(&msgInput,
		validation.Field(&msgInput.SenderID, validation.Required),
		validation.Field(&msgInput.ReceiverID, validation.Required),
		validation.Field(&msgInput.Message, validation.Required),
		validation.Field(&msgInput.Message, validation.Required),
	)
}
