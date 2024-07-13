package auth

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(rt *echo.Echo, db *gorm.DB) {
	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route := rt.Group("")
	{
		route.POST("/login", ctrl.SignIn)
		route.POST("/register", ctrl.Register)
	}
}
