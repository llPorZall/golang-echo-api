package middlewares

import (
	"api/pkg"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

//Authorization function
func Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get("Authorization")
		fmt.Println(header)
		jwtToken := strings.Fields(header)
		if len(jwtToken) != 2 {
			return c.JSON(http.StatusForbidden, pkg.ErrorResponse{Message: "You don't have permission", StatusCode: http.StatusForbidden})
		}
		email, err := pkg.VerifyJWTToken(jwtToken[1])
		if err != nil {
			return c.JSON(http.StatusForbidden, pkg.ErrorResponse{Message: "You don't have permission", StatusCode: http.StatusForbidden})
		}
		c.Request().Header.Set("email", email)
		return next(c)
	}
}
