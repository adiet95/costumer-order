package auth

import (
	"github.com/adiet95/go-order-api/src/database/models"
	"github.com/adiet95/go-order-api/src/interfaces"
	"github.com/adiet95/go-order-api/src/libs"
)

type auth_service struct {
	repo interfaces.AuthRepo
}
type token_response struct {
	Tokens string `json:"token"`
}

func NewService(reps interfaces.AuthRepo) *auth_service {
	return &auth_service{reps}
}

func (u auth_service) Login(body models.User) *libs.Response {
	checkRegist := libs.Validation(body.Email, body.Password)
	if checkRegist != nil {
		return libs.New(checkRegist.Error(), 400, true)
	}

	user, err := u.repo.FindByEmail(body.Email)
	if err != nil {
		return libs.New("email not registered, register first", 400, true)
	}
	if !libs.CheckPass(user.Password, body.Password) {
		return libs.New("wrong password", 400, true)
	}
	token := libs.NewToken(body.Email, user.Role)
	theToken, err := token.Create()
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	return libs.New(token_response{Tokens: theToken}, 200, false)
}

func (u auth_service) Register(body *models.User) *libs.Response {
	checkRegist := libs.Validation(body.Email, body.Password)
	if checkRegist != nil {
		return libs.New(checkRegist.Error(), 400, true)
	}

	hassPass, err := libs.HashPassword(body.Password)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	body.Password = hassPass
	body.Role = "user"
	result, err := u.repo.RegisterEmail(body)
	if err != nil {
		return libs.New("email already registered, please login", 400, true)
	}
	return libs.New(result, 200, false)
}
