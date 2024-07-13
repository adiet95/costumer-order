package middleware

import (
	"github.com/labstack/echo/v4"
	"strings"

	"github.com/adiet95/go-order-api/src/libs"
)

func CheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		headerToken := c.Request().Header.Get("Authorization")

		if !strings.Contains(headerToken, "Bearer") {
			return libs.New("invalid header type", 401, true).Send(c)

		}
		token := strings.Replace(headerToken, "Bearer ", "", -1)

		checkToken, err := libs.CheckToken(token)
		if err != nil {
			return libs.New(err.Error(), 401, true).Send(c)
		}

		c.Set("email", checkToken.Email)
		if err = next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}
