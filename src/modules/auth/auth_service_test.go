package auth

import (
	"testing"

	"github.com/adiet95/costumer-order/src/database/models"
	"github.com/adiet95/costumer-order/src/libs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewService(&repo)

	user, _ := libs.HashPassword("user12345678")
	var dataMock = models.User{Email: "user2@gmail.com", Password: "user12345678"}

	var dataMocks = models.User{Email: "user2@gmail.com", Password: user}

	repo.mock.On("RegisterEmail", &dataMock).Return(&dataMocks, nil)
	data := service.Register(&dataMock)
	res := data.Data.(*models.User)
	var expectEmail string = "user2@gmail.com"
	var expectPassword string = user

	assert.Equal(t, expectEmail, res.Email, "Email must be user")
	assert.Equal(t, expectPassword, res.Password, "Password must me hashing")

}
