package order

import (
	"errors"

	"github.com/adiet95/go-order-api/src/database/models"
	"gorm.io/gorm"
)

type order_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *order_repo {
	return &order_repo{db}
}

func (r *order_repo) FindAll(limit, offset int) (*models.Orders, error) {
	var datas *models.Orders
	result := r.db.Model(&datas).Limit(limit).Offset(offset).Find(&datas)
	if result.Error != nil {
		return nil, errors.New("failed obtain datas")
	}
	return datas, nil
}

func (r *order_repo) Save(data *models.Order) (*models.Order, error) {
	data.TotalPrice = data.Amount * data.Price

	res := r.db.Create(data)
	if res.Error != nil {
		return nil, errors.New("failed obtain datas")
	}
	return data, nil
}

func (re *order_repo) Update(data *models.Order, id int) (*models.Order, error) {
	data.TotalPrice = data.Amount * data.Price
	res := re.db.Model(&data).Where("order_id = ?", id).Updates(&data)

	if res.Error != nil {
		return nil, errors.New("failed to update data")
	}
	return data, nil
}

func (re *order_repo) Delete(id int) (*models.Order, error) {
	var data *models.Order
	var datas *models.Orders
	res := re.db.Where("order_id = ?", id).Find(&datas)

	if res.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	r := re.db.Model(data).Where("order_id = ?", id).Delete(&data)
	if r.Error != nil {
		return nil, errors.New("failed to delete data")
	}
	return nil, nil
}

func (re *order_repo) FindByName(name string) (*models.Orders, error) {
	var datas *models.Orders
	res := re.db.Order("order_id asc").Where("LOWER(order_name) LIKE ?", "%"+name+"%").Find(&datas)
	if res.Error != nil {
		return nil, errors.New("failed to found data")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	return datas, nil
}

func (re *order_repo) FindById(id int) (*models.Orders, error) {
	var datas *models.Orders
	res := re.db.Order("order_id asc").Where("order_id = ?", id).Find(&datas)
	if res.Error != nil {
		return nil, errors.New("failed to found data")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	return datas, nil
}

func (r *order_repo) GetUserId(email string) (*models.User, error) {
	var users *models.Users
	var user *models.User

	result := r.db.Model(&users).Where("email = ?", email).Find(&user)
	if result.Error != nil {
		return nil, errors.New("invalid user_id")
	}
	return user, nil
}
