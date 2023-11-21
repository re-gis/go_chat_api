package database

import "mime/multipart"

type User struct {
	ID       uint                  `json:"id" gorm:"primaryKey"`
	Email    string                `form:"email" json:"email" binding:"required"`
	Name     string                `form:"name" json:"name" binding:"required"`
	Password string                `form:"password" json:"password" binding:"required"`
	Profile  *multipart.FileHeader `form:"profile" json:"-" binding:"-"`
}
