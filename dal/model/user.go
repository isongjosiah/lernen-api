package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Student   bool   `json:"student"`
	//Enrolled  []string `gorm:"type:[]text"` TODO(josiah): sql doesn't allow a slice of strings. figure out a way around that
}

type LoginDetails struct {
	Email    string
	Password string
}

type UserDetails struct {
	User  *User
	Token string
}