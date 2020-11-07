package dal

import (
	"fmt"
	"net/http"

	"github.com/isongjosiah/lernen-api/dal/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
)

type IUserDAL interface {
	Add(user *model.User) (int, error)
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

// SetUp configures the DAL object
func (u *UserDAL) SetUp(db *gorm.DB) {
	u.Database = db
}

// NewUserDAL creates an instance of a user DAL
func NewUserDAL() *UserDAL {
	return &UserDAL{}
}

// Add creates a new User
func (u *UserDAL) Add(user *model.User) (int, error) {
	fmt.Println("DEBUG 4")
	db := u.Database
	fmt.Println("DEBUG 3")
	//check if user already exists in database.
	//TODO(josiah): check out what gorm.DB.NewRecord does.

	// check if email already exists in database
	account, _ := checkuser(u.Database, user.Email, user.Username)
	if account != nil {
		err := errors.New("User is already registered, please login")
		return http.StatusBadRequest, err
	}

	// Add the user here
	err := db.Debug().Create(user).Error
	if err != nil {
		err := errors.Wrap(err, "There was an error in adding user to the database")
		err = errors.New(err.Error())
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func checkuser(db *gorm.DB, email string, username string) (*model.User, error) {
	user := &model.User{}
	if err := db.Debug().Table("users").Where("email = ? OR username =?", email, username).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
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
	db := u.Database
	return checkuser(db, email, "")
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
