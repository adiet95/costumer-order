package order

import (
	"github.com/adiet95/costumer-order/src/middleware"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(rt *echo.Echo, db *gorm.DB) {
	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route := rt.Group("/order")
	route.Use(middleware.CheckAuth)
	{
		route.POST("", ctrl.Add)
		route.PUT("/:id", ctrl.Update)
		route.DELETE("/:id", ctrl.Delete)
		route.GET("", ctrl.GetAll)
		route.GET("/search", ctrl.Search)
		route.GET("/:id", ctrl.SearchId)
	}
}
