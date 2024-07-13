package interfaces

import (
	"github.com/adiet95/costumer-order/src/database/models"
	"github.com/adiet95/costumer-order/src/libs"
)

type AuthRepo interface {
	FindByEmail(username string) (*models.User, error)
	RegisterEmail(data *models.User) (*models.User, error)
}

type AuthService interface {
	Login(body models.User) *libs.Response
	Register(body *models.User) *libs.Response
}
