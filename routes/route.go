package routes

import (
	"Documents/belee/controllers"

	"github.com/labstack/echo/v4"
)

func NewRoutes() *echo.Echo {
	e := echo.New()
	ev1 := e.Group("api/v1/")
	ev1.GET("buyers", controllers.GetBuyersController)
	ev1.GET("buyers/login", controllers.LoginController)
	ev1.POST("buyers/register", controllers.RegisterController)
	ev1.GET("buyers/:buyersId", controllers.DetailsBuyers)
	return e
}
