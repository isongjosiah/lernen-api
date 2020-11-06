package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Firstname string `gorm:"type:char" json:"firstname"`
	Lastname  string `gorm:"type:char" json:"lastname"`
	Username  string `gorm:"type:char" json:"username"`
	Email     string `json:"email"`
	Password  string `gorm:"type:char" json:"password"`
	Student   bool
	//Enrolled  []string `gorm:"type:[]text"` TODO(josiah): sql doesn't allow a slice of strings. figure out a way around that
}
