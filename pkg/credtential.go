package pkg

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

//GenerateToken function
func GenerateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})
	tokenString, err := token.SignedString([]byte("MYKEYISSTRONG"))
	return tokenString, err
}

//VerifyJWTToken function
func VerifyJWTToken(token string) (string, error) {
	claims := &claims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("MYKEYISSTRONG"), nil
	})
	if err != nil {
		return "", err
	}
	return claims.Email, nil
}

//GeneratePassword function
func GeneratePassword(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	return string(bytes), err
}

//VerifyPassword function
func VerifyPassword(hash string, pwd string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd)); err != nil {
		return err
	}
	return nil
}
