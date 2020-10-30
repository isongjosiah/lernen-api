package dal

import "github.com/isongjosiah/lernen-api/dal/model"

type IUserDAL interface {
	Add(user *model.User) error
}

// UserDAL ...
type UserDAL struct {
}

// Setup configures the DAL object
func (u *UserDAL) Setup() error {

	return nil
}

// NewUserDAL creates an instance of a user DAL
func NewUserDAL() *UserDAL {
	return &UserDAL{}
}

// Add creates a new User
func (u UserDAL) Add(user *model.User) error {

	return nil
}
