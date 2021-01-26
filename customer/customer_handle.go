package customer

import (
	"api/pkg"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type handle struct {
	repository Repository
}

// Customer registor api
// @Summary Customer registor.
// @Description Customer registor api.
// @Tags customer
// @Produce  json
// @Param age formData int true "Age"
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Param firstname formData string true "FirstName"
// @Param lastname formData string true "LastName"
// @Success 200 {object} pkg.SuccessResponse
// @Failure 400,404,500 {object} pkg.ErrorResponse
// @Router /customer [post]
func (h *handle) customerRegistor(c echo.Context) error {
	request := CustomerRegistorBody{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "Error please try again", StatusCode: http.StatusBadRequest})
	}
	fmt.Println(request)
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return c.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Missing parameter", StatusCode: http.StatusNotFound})
	}
	if existingCustomer, _ := h.repository.FindOne(bson.M{"email": request.Email}); existingCustomer != nil {
		return c.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Email not available", StatusCode: http.StatusNotFound})
	}

	password, err := pkg.GeneratePassword(request.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: "Server error", StatusCode: http.StatusInternalServerError})
	}
	request.Password = password
	customer, err := h.repository.Create(request)
	if err != nil {
		return c.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Can't create customer", StatusCode: http.StatusInternalServerError})
	}

	return c.JSON(http.StatusOK, pkg.SuccessResponse{
		Message:    fmt.Sprintf("Create customer %s success with customer id %s", customer.Email, customer.Id.Hex()),
		StatusCode: http.StatusOK,
		Data: CustomerResponseBody{
			Email:     customer.Email,
			FirstName: customer.FirstName,
			Lastname:  customer.Lastname,
			Age:       customer.Age,
		},
	})
}

// Customer Login api
// @Summary Customer Login.
// @Description Customer Login api.
// @Tags customer
// @Produce  json
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Success 200 {object} pkg.SuccessResponse
// @Failure 400,404,500 {object} pkg.ErrorResponse
// @Router /customer/login [post]
func (h *handle) customerLogin(c echo.Context) error {
	request := CustomerLoginBody{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "Error please try again", StatusCode: http.StatusBadRequest})
	}
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return c.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Missing parameter", StatusCode: http.StatusNotFound})
	}
	customer, err := h.repository.FindOne(bson.M{"email": request.Email})
	if err != nil {
		return c.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Customer not found", StatusCode: http.StatusNotFound})
	}
	if err = pkg.VerifyPassword(customer.Password, request.Password); err != nil {
		return c.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Email or Password not match", StatusCode: http.StatusNotFound})
	}

	jwtToken, err := pkg.GenerateToken(customer.Email)
	if err != nil {
		return c.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Email or Password not match", StatusCode: http.StatusNotFound})
	}

	return c.JSON(http.StatusOK, pkg.SuccessResponse{
		Message:    "Login success",
		StatusCode: http.StatusOK,
		Data:       jwtToken,
	})
}

// Customer Get Information
// @Summary Customer information.
// @Description Customer get information api.
// @Tags customer
// @Produce  json
// @Param Authorization header string true "Bearer"
// @Success 200 {object} pkg.SuccessResponse
// @Failure 400,404,500 {object} pkg.ErrorResponse
// @Router /customer/me [get]
func (h *handle) customerInformation(c echo.Context) error {
	customer, err := h.repository.FindOne(bson.M{"email": c.Request().Header.Get("email")})
	if err != nil {
		return c.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Customer not found", StatusCode: http.StatusNotFound})
	}
	cusRes := CustomerResponseBody{
		Email:     customer.Email,
		FirstName: customer.FirstName,
		Lastname:  customer.Lastname,
		Age:       customer.Age,
	}
	return c.JSON(http.StatusOK, pkg.SuccessResponse{
		StatusCode: http.StatusOK,
		Message:    "Get customer information success",
		Data:       cusRes,
	})
}

// Customer change password api
// @Summary Customer change password.
// @Description Customer change password api.
// @Tags customer
// @Produce  json
// @Param Authorization header string true "Bearer"
// @Param oldPassword formData string true "OldPassword"
// @Param newPassword formData string true "NewPassword"
// @Success 200 {object} pkg.SuccessResponse
// @Failure 400,404,500 {object} pkg.ErrorResponse
// @Router /customer/password [patch]
func (h *handle) customerChangePassword(c echo.Context) error {
	request := CustomerChangePassword{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "Error please try again", StatusCode: http.StatusBadRequest})
	}
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return c.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Missing parameter", StatusCode: http.StatusNotFound})
	}

	customer, err := h.repository.FindOne(bson.M{"email": c.Request().Header.Get("email")})
	if err != nil {
		return c.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Customer not found", StatusCode: http.StatusNotFound})
	}

	if err = pkg.VerifyPassword(customer.Password, request.OldPassword); err != nil {
		return c.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Email or Password not match", StatusCode: http.StatusNotFound})
	}

	password, err := pkg.GeneratePassword(request.NewPassword)
	if err != nil {
		return c.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Can't update customer password", StatusCode: http.StatusNotFound})
	}

	fmt.Println(password)
	err = h.repository.Update(bson.M{"email": c.Request().Header.Get("email")}, bson.M{"$set": bson.M{"password": password}})
	if err != nil {
		return c.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Can't update customer password", StatusCode: http.StatusNotFound})
	}

	return c.JSON(http.StatusOK, pkg.SuccessResponse{
		Message:    "Customer chnage password success",
		StatusCode: http.StatusOK,
	})
}
