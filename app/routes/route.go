package routes

import (
	"final_project/belee/controllers/buyers"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	BuyerController buyers.BuyerController
}

func (c1 *ControllerList) RouteRegister(e *echo.Echo) {
	e.POST("buyers/login", c1.BuyerController.Login)
}
