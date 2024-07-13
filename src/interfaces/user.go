package interfaces

import (
	"github.com/adiet95/costumer-order/src/database/models"
	"github.com/adiet95/costumer-order/src/libs"
)

type UserRepo interface {
	FindAll(limit, offset int) (*models.Users, error)
	Save(data *models.User) (*models.User, error)
	UpdateUser(data *models.User, email string) (*models.User, error)
	DeleteUser(id int) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByName(name string) (*models.Users, error)
	FindById(id int) (*models.User, error)
}

type UserService interface {
	Add(data *models.User) *libs.Response
	Update(data *models.User, email, emailUpdate string) *libs.Response
	Delete(id int) *libs.Response
	FindEmail(email string) *libs.Response
	FindEmails(email string, limit, offset int) *libs.Response
	Search(email string) *libs.Response
	SearchName(name string) *libs.Response
	GetById(id int) *libs.Response
}
