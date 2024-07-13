package auth

import (
	"errors"

	"github.com/adiet95/costumer-order/src/database/models"
	"gorm.io/gorm"
)

type auth_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *auth_repo {
	return &auth_repo{db}
}

func (re *auth_repo) FindByEmail(email string) (*models.User, error) {
	var data *models.User
	var datas *models.Users

	res := re.db.Model(&datas).Where("email = ?", email).Find(&data)
	if res.Error != nil {
		return nil, errors.New("failed to find data")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("email not found")
	}
	return data, nil
}

func (re *auth_repo) RegisterEmail(data *models.User) (*models.User, error) {
	var datas *models.Users

	res := re.db.Model(&datas).Where("email = ?", data.Email).Find(&data)
	if res.Error != nil {
		return nil, errors.New("failed to find data")
	}
	if res.RowsAffected > 0 {
		return nil, errors.New("email registered, go to login")
	}

	r := re.db.Create(data)
	if r.Error != nil {
		return nil, errors.New("failled to obtain data")
	}
	return data, nil
}
