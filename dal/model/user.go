package model

type User struct {
	Firstname string `gorm:"type:char" json:"firstname"`
	Lastname  string `gorm:"type:car" json:"lastname"`
	Username  string `gorm:"type:char" json:"username"`
	Email     string `json:"email"`
	Password  string `gorm:"type:char" json:"password"`
	Student   bool
	Enrolled  []string `gorm:"type:[]text"`
}
