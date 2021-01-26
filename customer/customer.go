package customer

import (
	"api/db"
	"api/middlewares"

	"github.com/labstack/echo/v4"
)

//CustomerAPI regis resource
func CustomerAPI(app *echo.Group, resource *db.Resource) {
	repository := newCustomerRepository(resource)
	handle := handle{repository: repository}
	app.POST("/customer", handle.customerRegistor)
	app.POST("/customer/login", handle.customerLogin)
	app.GET("/customer/me", handle.customerInformation, middlewares.Authorization)
	app.PATCH("/customer/password", handle.customerChangePassword, middlewares.Authorization)
}
