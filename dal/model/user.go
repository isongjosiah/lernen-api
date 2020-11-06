package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Firstname string `gorm:"type:varchar" json:"firstname"`
	Lastname  string `gorm:"type:varchar" json:"lastname"`
	Username  string `gorm:"type:varchar" json:"username"`
	Email     string `json:"email"`
	Password  string `gorm:"type:varchar" json:"password"`
	Student   bool
	//Enrolled  []string `gorm:"type:[]text"` TODO(josiah): sql doesn't allow a slice of strings. figure out a way around that
}
