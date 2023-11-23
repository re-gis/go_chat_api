package database

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Profile  string `json:"profile" form:"profile"`
}

type Message struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Body       string `json:"body"`
	ChatId     uint
	SenderId   uint `json:"sender"`
	ReceiverId uint `json:"receiver"`
	Isread     bool
}

type Chat struct {
	ID         uint `json:"id" gorm:"primaryKey"`
	SenderId   uint
	ReceiverId uint
}
