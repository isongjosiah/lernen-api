package dal

import (
	"github.com/isongjosiah/lernen-api/dal/model"
	"github.com/jinzhu/gorm"
)

type ICoursesDAL interface {
	Create(course *model.Course) error
	Delete(title string) error
	Fetch(id int) (*model.Course, error)
}

// CourseDAL ...
type CourseDAL struct {
	Database *gorm.DB
}

// NewCourseDAL creates an instance of a user DAL
func NewCourseDAL(db *gorm.DB) *CourseDAL {
	return &CourseDAL{
		Database: db,
	}
}

//Create creates a new course that would be scheduled for review
func (c *CourseDAL) Create(course *model.Course) error {
	return nil
}

//Delete deletes a course based on the title
func (c *CourseDAL) Delete(title string) error {
	return nil
}

//Fetch returns the details of a course based on the id
func (c *CourseDAL) Fetch(id int) error {
	return nil
}
