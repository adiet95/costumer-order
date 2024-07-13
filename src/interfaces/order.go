package interfaces

import (
	"github.com/adiet95/go-order-api/src/database/models"
	"github.com/adiet95/go-order-api/src/libs"
)

type OrderRepo interface {
	FindAll(limit, offset int) (*models.Orders, error)
	Save(data *models.Order) (*models.Order, error)
	Update(data *models.Order, id int) (*models.Order, error)
	Delete(id int) (*models.Order, error)
	FindByName(name string) (*models.Orders, error)
	GetUserId(email string) (*models.User, error)
	FindById(id int) (*models.Orders, error)
}

type OrderService interface {
	GetAll(limit, offset int) *libs.Response
	Add(data *models.Order, email string) *libs.Response
	Update(data *models.Order, id int, email string) *libs.Response
	Delete(id int) *libs.Response
	Search(name string) *libs.Response
	SearchId(id int) *libs.Response
}
