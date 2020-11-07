package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/isongjosiah/lernen-api/dal/model"
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
	//read the information from the body. TODO(josiah): check if you need to define middlewares to set the content-type to "application/json"
	err := decodeJSONBody(nil, r.Body, &user)
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// check for empty fields
	if user.Firstname == "" || user.Lastname == "" || user.Email == "" || user.Username == "" || user.Password == "" {
		WriteErrorResponse(w, http.StatusBadRequest, "some required fields are empty. Please fill all fields")
		return
	}
	fmt.Println("DEBUG1")
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	user.Password = hashPassword(user.Password, w)
	fmt.Println("DEBUG2")
	fmt.Printf("%T", &user)
	// add the user to the database
	status, err := a.Deps.DAL.UserDAL.Add(&user)
	if err != nil {
		WriteErrorResponse(w, status, err.Error())
		return
	}

	WriteJSONPayload(w, &ServerResponse{
		Message:    "User successfully registered",
		StatusCode: status,
		Payload:    user,
	})
}

type LoginDetails struct {
	Email    string
	Password string
}

type UserDetails struct {
	User  *model.User
	Token string
}

//Login is the handler for the path /login
func (a *API) Login(w http.ResponseWriter, r *http.Request) {
	var user *model.User
	var loginDetails LoginDetails
	var userDetails UserDetails

	err := decodeJSONBody(nil, r.Body, &loginDetails)

	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if len(loginDetails.Email) == 0 || len(loginDetails.Password) == 0 {
		WriteErrorResponse(w, http.StatusBadRequest, "some required fields are empty. Please fill all fields")
		return
	}

	//Find user by email
	user, findUserErr := a.Deps.DAL.UserDAL.FindUserByEmail(loginDetails.Email)
	fmt.Println("Retrieving user details")

	if findUserErr == nil {
		if !comparePasswords(user.Password, []byte(loginDetails.Password)) {
			WriteErrorResponse(w, http.StatusBadRequest, "User details do not match.")
			return
		}

		var jwtSecretKey = []byte("jwt_secret_key")
		tokenString, tokenErr := GenerateToken(jwtSecretKey, loginDetails.Email)

		if tokenErr == nil {
			userDetails.Token = tokenString
			userDetails.User = user
			WriteJSONPayload(w, &ServerResponse{
				Message:    "Login successful",
				StatusCode: 200,
				Payload:    userDetails,
			})
			return
		}
		WriteJSONPayload(w, &ServerResponse{
			Message:    "Error while validating user",
			StatusCode: http.StatusInternalServerError,
			Payload:    userDetails,
		})
		return
	}
	WriteErrorResponse(w, http.StatusBadRequest, "User not found.")
	return
}

func hashPassword(password string, w http.ResponseWriter) string {
	//hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	fmt.Println("DEBUG1")
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
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
