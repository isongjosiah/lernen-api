package dal

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/badoux/checkmail"

	"github.com/isongjosiah/lernen-api/dal/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
	"net/http"
)

type IUserDAL interface {
	Add(user *model.User) (int, error)
	Delete(username string) error
	FindUserByUsername(username string) (*model.User, error)
	FindUserByEmail(email string) (*model.User, error)
	GetCourses(user *model.User) (*[]string, error)
	EditPicture(username string, data *s3manager.UploadOutput) error
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
	db := u.Database

	// check if the email provided is valid
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return http.StatusBadRequest, errors.New("invalid email format. Please use \"username@example.com\" as format.")
	}

	// check if email already exists in database
	account, _ := checkUser(db, "email", user.Email)
	if account != nil {
		err := errors.New("User is already registered, please login")
		return http.StatusBadRequest, err
	}

	// check if username already exists in database
	account, _ = checkUser(db, "username", user.Username)
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

func checkUser(db *gorm.DB, field string, input string) (*model.User, error) {
	user := &model.User{}
	query := fmt.Sprintf("%s = ?", field)
	if err := db.Debug().Table("users").Where(query, input).First(user).Error; err != nil {
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
	db := u.Database
	return checkUser(db, "username", username)
}

// FindUserByEmail returns a user based on a provided email address
func (u *UserDAL) FindUserByEmail(email string) (*model.User, error) {
	db := u.Database
	return checkUser(db, "email", email)
}

// GetCourses returns a list of the title of the courses the user is enrolled in
func (u *UserDAL) GetCourses(user *model.User) (*[]string, error) {
	var courses []string
	return &courses, nil
}

// EditProfile edits the profile of a user
func (u *UserDAL) EditPicture(username string, data *s3manager.UploadOutput) error {
	db := u.Database
	user, err := u.FindUserByUsername(username)
	if err != nil {
		return err
	}
	user.PicSrc = data.Location
	user.PicUploadID = data.UploadID
	user.PicVersionID = data.VersionID

	db.Save(user)
	return nil
}
