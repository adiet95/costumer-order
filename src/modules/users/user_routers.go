package users

import (
	"github.com/adiet95/costumer-order/src/middleware"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(rt *echo.Echo, db *gorm.DB) {
	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route := rt.Group("/user")
	route.Use(middleware.CheckAuth)
	{
		route.GET("", ctrl.GetAll)
		route.GET("/:id", ctrl.SearchId)
		route.POST("", ctrl.Add, middleware.CheckAuthor)
		route.PUT("", ctrl.Update)
		route.DELETE("/:id", ctrl.Delete, middleware.CheckAuthor)
		route.GET("/detail", ctrl.Search, middleware.CheckAuthor)
		route.GET("/search", ctrl.SearchName, middleware.CheckAuthor)
	}
}
