package model

type User struct {
	Firstname string
	Lastname  string
	Username  string
	Password  string
	Student   bool
	Enrolled  []string `gorm:"type:[]text"` // This contains the tile of the courses that user is enrolled in
}
