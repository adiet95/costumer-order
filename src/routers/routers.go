package routers

import (
	"errors"
	"github.com/labstack/echo/v4"

	"github.com/adiet95/costumer-order/src/database"
	auth "github.com/adiet95/costumer-order/src/modules/auth"
	"github.com/adiet95/costumer-order/src/modules/order"
	"github.com/adiet95/costumer-order/src/modules/users"
)

func New(mainRoute *echo.Echo) (*echo.Echo, error) {
	db, err := database.New()
	if err != nil {
		return nil, errors.New("failed init database")
	}

	auth.New(mainRoute, db)
	users.New(mainRoute, db)
	order.New(mainRoute, db)

	return mainRoute, nil
}
