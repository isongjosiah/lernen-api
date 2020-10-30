package model

type User struct {
	Firstname string
	Lastname  string
	Username  string
	Password  string
	Enrolled  *[]Course
}
