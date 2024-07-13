package order

import (
	"github.com/adiet95/costumer-order/src/database/models"
	"github.com/adiet95/costumer-order/src/interfaces"
	"github.com/adiet95/costumer-order/src/libs"
)

type order_service struct {
	order_repo interfaces.OrderRepo
}

func NewService(reps interfaces.OrderRepo) *order_service {
	return &order_service{reps}
}

func (r *order_service) GetAll(limit, offset int) *libs.Response {
	data, err := r.order_repo.FindAll(limit, offset)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (re *order_service) Add(data *models.Order, email string) *libs.Response {
	findId, err := re.order_repo.GetUserId(email)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	data.UserId = int(findId.UserId)
	data.User.Email = findId.Email
	data.User.FullName = findId.FullName

	result, err := re.order_repo.Save(data)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(result, 200, false)
}

func (re *order_service) Update(data *models.Order, id int, email string) *libs.Response {
	findId, err := re.order_repo.GetUserId(email)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	data.UserId = int(findId.UserId)
	data.User.Email = findId.Email
	data.User.FullName = findId.FullName

	res, err := re.order_repo.Update(data, id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(res, 200, false)
}

func (re *order_service) Delete(id int) *libs.Response {
	data, err := re.order_repo.Delete(id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (re *order_service) Search(name string) *libs.Response {
	data, err := re.order_repo.FindByName(name)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (re *order_service) SearchId(id int) *libs.Response {
	data, err := re.order_repo.FindById(id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}
