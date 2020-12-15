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

type AuthResponse struct {
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}
type RegistrationDetail struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginDetails struct {
	Email       string
	Password    string `json:"password"`
	SocialLogin bool   `json: "socialLogin"`
}

type UserDetails struct {
	UserInfo *AuthResponse `json:"user"`
	Token    string        `json:"token"`
}
