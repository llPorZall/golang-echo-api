package customer

import (
	"api/pkg"

	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyPassword(t *testing.T) {
	hashPassword, _ := pkg.GeneratePassword("123456")
	err := pkg.VerifyPassword(hashPassword, "123456")
	fmt.Println(err)
	assert.Equal(t, err, nil, "Error should be nil")
}

func TestVerifyJWT(t *testing.T) {
	email := "pongchai@neversitup.com"
	jwtToken, _ := pkg.GenerateToken(email)
	jwtTokenDecode, _ := pkg.VerifyJWTToken(jwtToken)
	assert.Equal(t, email, jwtTokenDecode, "Email should be same")
}
func TestRegistorCustomer(t *testing.T) {
	//TODO
}
func TestUpdateCustomer(t *testing.T) {
	//TODO
}

func TestChangeCustomerPassword(t *testing.T) {
	//TODO
}
