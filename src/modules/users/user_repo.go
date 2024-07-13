package users

import (
	"errors"

	"github.com/adiet95/costumer-order/src/database/models"
	"gorm.io/gorm"
)

type user_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *user_repo {
	return &user_repo{db}
}

func (r *user_repo) FindAll(limit, offset int) (*models.Users, error) {
	var data models.Users

	result := r.db.Model(&data).Limit(limit).Offset(offset).Order("email asc").Find(&data)

	if result.Error != nil {
		return nil, errors.New("failed obtain datas")
	}
	return &data, nil
}

func (r *user_repo) Save(data *models.User) (*models.User, error) {
	var datas models.Users
	res := r.db.Where("LOWER(email) = ?", data.Email).Find(&datas)

	if res.RowsAffected != 0 {
		return nil, errors.New("email already registered")
	}

	result := r.db.Create(data)
	if result.Error != nil {
		return nil, errors.New("failled to obtain data")
	}
	return data, nil
}

func (re *user_repo) UpdateUser(data *models.User, email string) (*models.User, error) {
	res := re.db.Model(&data).Where("LOWER(email) = ?", email).Updates(&data)

	if res.Error != nil {
		return nil, errors.New("failed to update data")
	}
	return data, nil
}

func (re *user_repo) DeleteUser(id int) (*models.User, error) {
	var data *models.User
	var datas *models.Users
	res := re.db.Where("user_id = ?", id).Find(&datas)

	if res.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	r := re.db.Model(data).Where("user_id = ?", id).Delete(&data)
	if r.Error != nil {
		return nil, errors.New("failed to delete data")
	}
	return nil, nil
}

func (re *user_repo) FindByEmails(email string, limit, offset int) (*models.Users, error) {
	//var data *models.User
	var datas *models.Users

	res := re.db.Model(&datas).Where("LOWER(email) = ?", email).Limit(limit).Offset(offset).Order("email asc").Find(&datas)
	if res.Error != nil {
		return nil, errors.New("failed to find data")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("email not found")
	}
	return datas, nil
}
func (re *user_repo) FindByEmail(email string) (*models.User, error) {
	var data *models.User
	res := re.db.Model(&data).Where("LOWER(email) = ?", email).First(&data)
	if res.Error != nil {
		return nil, errors.New("failed to find data")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("email not found")
	}
	return data, nil
}

func (re *user_repo) FindByName(name string) (*models.Users, error) {
	var datas *models.Users

	res := re.db.Model(&datas).Where("LOWER(full_name) LIKE ?", "%"+name+"%").Find(&datas)
	if res.Error != nil {
		return nil, errors.New("failed to find data")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("name not found")
	}
	return datas, nil
}

func (re *user_repo) FindById(id int) (*models.User, error) {
	var datas *models.User

	res := re.db.Model(&datas).Where("user_id = ?", id).First(&datas)
	if res.Error != nil {
		return nil, errors.New("failed to find data")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("name not found")
	}
	return datas, nil
}
