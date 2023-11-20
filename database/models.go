package database

type User struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Profile string `json:"profile"`
}

