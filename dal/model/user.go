package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Firstname    string  `json:"firstname"`
	Lastname     string  `json:"lastname"`
	Username     string  `json:"username"`
	Email        string  `json:"email"`
	Password     string  `json:"password"`
	Student      bool    `json:"student"`
	PicSrc       string  `json:"pic_src"`
	PicVersionID *string `json:"pic_version_id"`
	PicUploadID  string  `json:"pic_upload_id"`
}

type AuthResponse struct {
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	PicSrc    string `json:"pic_src"`
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
