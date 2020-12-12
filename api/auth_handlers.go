package api

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/isongjosiah/lernen-api/dal/model"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

//AuthRoutes sets up the authentication handlers
func (a *API) AuthRoutes(router *chi.Mux) http.Handler {
	router.Method("POST", "/login", http.HandlerFunc(a.Login))
	router.Method("POST", "/register", http.HandlerFunc(a.Register))

	return router
}

//Register is the handler for the path /register
func (a *API) Register(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var registrationDetails model.RegistrationDetail
	//read the information from the body. TODO(josiah): check if you need to define middlewares to set the content-type to "application/json"
	err := decodeJSONBody(nil, r.Body, &registrationDetails)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// check for empty fields
	if registrationDetails.Firstname == "" || registrationDetails.Lastname == "" || registrationDetails.Email == "" || registrationDetails.Username == "" || registrationDetails.Password == "" {
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("some required fields are empty. Please fill all fields"))
		return
	}

	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	user.Firstname = registrationDetails.Firstname
	user.Lastname =registrationDetails.Lastname
	user.Username = registrationDetails.Username
	user.Email = registrationDetails.Email
	user.Password = hashPassword(registrationDetails.Password, w)
	// add the user to the database
	status, err := a.Deps.DAL.UserDAL.Add(&user)
	if err != nil {
		WriteErrorResponse(w, status, err)
		return
	}

	res := &model.AuthResponse{
		Model: gorm.Model{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		},
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Username:  user.Username,
		Email:     user.Email,
	}

	WriteJSONPayload(w, &ServerResponse{
		Message:    "User successfully registered",
		StatusCode: status,
		Payload:    res,
	})
}

//Login is the handler for the path /login
func (a *API) Login(w http.ResponseWriter, r *http.Request) {
	var user *model.User
	var loginDetails model.LoginDetails
	var userDetails model.UserDetails

	err := decodeJSONBody(nil, r.Body, &loginDetails)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if loginDetails.Email == "" || loginDetails.Password == "" {
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("some required fields are empty. Please fill all fields"))
		return
	}
	// check if it is a valid email
	err = checkmail.ValidateFormat(loginDetails.Email)

	// If error is not equals to nil, then it must be a username
	if err != nil {
		user, err = a.Deps.DAL.UserDAL.FindUserByUsername(loginDetails.Email)
	} else {
		//Find user by email
		user, err = a.Deps.DAL.UserDAL.FindUserByEmail(loginDetails.Email)
		log.Info("Retrieving user details")
	}

	if err == nil {
		if !comparePasswords(user.Password, []byte(loginDetails.Password)) {
			WriteErrorResponse(w, http.StatusBadRequest, errors.New("user details do not match"))
			return
		}

		jwtSecretKey := []byte(a.Config.TokenSecret)
		tokenString, tokenErr := GenerateToken(jwtSecretKey, loginDetails.Email)

		if tokenErr != nil {
			WriteJSONPayload(w, &ServerResponse{
				Err:        "Error while validating user",
				Message:    "failed",
				StatusCode: http.StatusInternalServerError,
				Payload:    loginDetails,
			})
			return
		}

		res := &model.AuthResponse{
			Model: gorm.Model{
				ID:        user.ID,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
				DeletedAt: user.DeletedAt,
			},
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Username:  user.Username,
			Email:     user.Email,
		}
		userDetails.Token = tokenString
		userDetails.UserInfo = res
		WriteJSONPayload(w, &ServerResponse{
			Message:    "Login successful",
			StatusCode: http.StatusOK,
			Payload:    userDetails,
		})
		return
	}

	WriteErrorResponse(w, http.StatusBadRequest, errors.New("user not found"))
}

func hashPassword(password string, w http.ResponseWriter) string {
	//hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, err)
		return "Error while hashing password"
	}

	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
