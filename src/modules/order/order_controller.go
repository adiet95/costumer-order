package order

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"

	"github.com/adiet95/costumer-order/src/database/models"
	"github.com/adiet95/costumer-order/src/interfaces"
	"github.com/adiet95/costumer-order/src/libs"
)

type order_ctrl struct {
	svc interfaces.OrderService
}

func NewCtrl(reps interfaces.OrderService) *order_ctrl {
	return &order_ctrl{svc: reps}
}

func (re *order_ctrl) GetAll(c echo.Context) error {
	v := c.QueryParam("limit")
	limit, _ := strconv.Atoi(v)
	if limit == 0 {
		limit = 10
	}

	val := c.QueryParam("offset")
	offset, _ := strconv.Atoi(val)

	return re.svc.GetAll(limit, offset).Send(c)
}

func (re *order_ctrl) Add(c echo.Context) error {
	claim_user := c.Get("email")
	if claim_user == "" {
		return libs.New("claim user is not exist", 400, true).Send(c)
	}

	var data models.Order
	err := json.NewDecoder(c.Request().Body).Decode(&data)
	if err != nil {
		return libs.New(err.Error(), 400, true).Send(c)
	}
	return re.svc.Add(&data, claim_user.(string)).Send(c)
}

func (re *order_ctrl) Update(c echo.Context) error {
	claim_user := c.Get("email")
	if claim_user == "" {
		return libs.New("claim user is not exist", 400, true).Send(c)
	}
	email := claim_user.(string)
	val := c.QueryParam("id")
	v, _ := strconv.Atoi(val)

	var datas models.Order
	err := json.NewDecoder(c.Request().Body).Decode(&datas)
	if err != nil {
		return libs.New(err.Error(), 400, true).Send(c)
	}
	return re.svc.Update(&datas, v, email).Send(c)
}

func (re *order_ctrl) Delete(c echo.Context) error {
	val := c.Param("id")
	v, _ := strconv.Atoi(val)

	return re.svc.Delete(v).Send(c)
}

func (re *order_ctrl) Search(c echo.Context) error {
	val := c.QueryParam("name")
	v := strings.ToLower(val)
	return re.svc.Search(v).Send(c)
}

func (re *order_ctrl) SearchId(c echo.Context) error {
	val := c.Param("id")
	v, _ := strconv.Atoi(val)
	return re.svc.SearchId(v).Send(c)
}
