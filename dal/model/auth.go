package model

import "github.com/jinzhu/gorm"

type AuthResponse struct {
	gorm.Model
	Fullname string `json:"fullname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

type LoginDetails struct {
	Email    string
	Password string `json:"password"`
}

type UserDetails struct {
	UserInfo *AuthResponse `json:"user"`
	Token    string        `json:"token"`
}

type RegistrationDetails struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
