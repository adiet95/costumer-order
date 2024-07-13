package users

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"

	"github.com/adiet95/costumer-order/src/database/models"
	"github.com/adiet95/costumer-order/src/interfaces"
	"github.com/adiet95/costumer-order/src/libs"
)

type user_ctrl struct {
	svc interfaces.UserService
}

func NewCtrl(reps interfaces.UserService) *user_ctrl {
	return &user_ctrl{svc: reps}
}

func (re *user_ctrl) GetAll(c echo.Context) error {
	claim_user := c.Get("email")
	if claim_user == "" {
		libs.New("claim user is not exist", 400, true)
	}

	v := c.QueryParam("limit")
	limit, _ := strconv.Atoi(v)
	if limit == 0 {
		limit = 10
	}

	val := c.QueryParam("offset")
	offset, _ := strconv.Atoi(val)

	result := re.svc.FindEmails(claim_user.(string), limit, offset)
	return result.Send(c)
}

func (re *user_ctrl) Add(c echo.Context) error {
	var data models.User
	err := json.NewDecoder(c.Request().Body).Decode(&data)
	if err != nil {
		return libs.New(err.Error(), 400, true).Send(c)
	}
	return re.svc.Add(&data).Send(c)
}

func (re *user_ctrl) Update(c echo.Context) error {
	claim_user := c.Get("email")
	if claim_user == "" {
		return libs.New("claim user is not exist", 400, true).Send(c)
	}

	email := c.QueryParam("email")

	var data models.User
	err := json.NewDecoder(c.Request().Body).Decode(&data)
	if err != nil {
		return libs.New(err.Error(), 400, true).Send(c)
	}

	return re.svc.Update(&data, claim_user.(string), email).Send(c)
}

func (re *user_ctrl) Delete(c echo.Context) error {
	val := c.Param("id")
	id, _ := strconv.Atoi(val)
	return re.svc.Delete(id).Send(c)
}

func (re *user_ctrl) Search(c echo.Context) error {
	val := c.QueryParam("email")
	v := strings.ToLower(val)
	return re.svc.Search(v).Send(c)
}

func (re *user_ctrl) SearchName(c echo.Context) error {
	val := c.QueryParam("name")
	v := strings.ToLower(val)
	return re.svc.SearchName(v).Send(c)
}

func (re *user_ctrl) SearchId(c echo.Context) error {
	val := c.Param("id")
	id, _ := strconv.Atoi(val)
	return re.svc.GetById(id).Send(c)
}
