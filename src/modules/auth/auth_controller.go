package auth

import (
	"encoding/json"
	"github.com/labstack/echo/v4"

	"github.com/adiet95/costumer-order/src/database/models"
	"github.com/adiet95/costumer-order/src/interfaces"
	"github.com/adiet95/costumer-order/src/libs"
)

type user_ctrl struct {
	repo interfaces.AuthService
}

func NewCtrl(reps interfaces.AuthService) *user_ctrl {
	return &user_ctrl{reps}
}

func (u user_ctrl) SignIn(c echo.Context) error {
	var data models.User

	err := json.NewDecoder(c.Request().Body).Decode(&data)
	if err != nil {
		return libs.New(err.Error(), 401, true).Send(c)
	}

	return u.repo.Login(data).Send(c)
}

func (u user_ctrl) Register(c echo.Context) error {
	var data *models.User

	err := json.NewDecoder(c.Request().Body).Decode(&data)
	if err != nil {
		return libs.New(err.Error(), 401, true).Send(c)
	}
	return u.repo.Register(data).Send(c)
}
