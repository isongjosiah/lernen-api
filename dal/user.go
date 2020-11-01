package dal

import (
	"github.com/isongjosiah/lernen-api/dal/model"
	"github.com/jinzhu/gorm"
)

type IUserDAL interface {
	Add(user *model.User) error
	Delete(username string) error
	FindUserByUsername(username string) (*model.User, error)
	FindUserByEmail(email string) (*model.User, error)
	GetCourses(user *model.User) (*[]string, error)
	EditProfile(user *model.User) error
}

// UserDAL ...
type UserDAL struct {
	Database *gorm.DB
}

// NewUserDAL creates an instance of a user DAL
func NewUserDAL(db *gorm.DB) *UserDAL {
	return &UserDAL{
		Database: db,
	}
}

// Add creates a new User
func (u *UserDAL) Add(user *model.User) error {
	err := u.Database.Debug().Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a user
func (u *UserDAL) Delete(username string) error {
	return nil
}

// FindUserByUsername returns a user based on a provided username
func (u *UserDAL) FindUserByUsername(username string) (*model.User, error) {
	return nil, nil
}

// FindUserByEmail returns a user based on a provided email address
func (u *UserDAL) FindUserByEmail(email string) (*model.User, error) {
	return nil, nil
}

// GetCourses returns a list of the title of the courses the user is enrolled in
func (u *UserDAL) GetCourses(user *model.User) (*[]string, error) {
	var courses []string
	return &courses, nil
}

// EditProfile edits the profile of a user
func (u *UserDAL) EditProfile(user *model.User) error {
	return nil
}
