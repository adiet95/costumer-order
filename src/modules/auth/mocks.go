package auth

import (
	"github.com/adiet95/go-order-api/src/database/models"
	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	mock mock.Mock
}

func (m *RepoMock) FindByEmail(email string) (*models.User, error) {
	args := m.mock.Called(email)
	return args.Get(0).(*models.User), nil
}

func (m *RepoMock) RegisterEmail(data *models.User) (*models.User, error) {
	args := m.mock.Called(data)
	return args.Get(0).(*models.User), nil
}
