package api

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claim struct {
	Email    string
	Username string
	jwt.StandardClaims
}

// GenerateToken generates the token
func GenerateToken(jwtSecretKey string, email string) (string, error) {
	IssuedAt := time.Now()
	ExpiresAt := time.Now().Add(time.Hour * 24)
	claims := &Claim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: ExpiresAt.Unix(),
			IssuedAt:  IssuedAt.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}

// VerifyToken verifies the token sent alongside a request
func VerifyToken(jwtSecret string, tokenString string) (string, error) {
	claim := &Claim{}
	//TODO(josiah): find out how you can verify the algorithm used before validating token.
	token, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if token != nil {
		return claim.Email, nil
	}

	return "", err
}
