package model

import "github.com/jinzhu/gorm"

type Course struct {
	gorm.Model
	Title    string
	Content  string
	Progress int
}
