package database

type User struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Profile string `json:"profile" form:"profile"`
}

