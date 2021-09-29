package routes

import (
	"final_project/belee/app/middleware"
	"final_project/belee/controllers/buyers"
	"final_project/belee/controllers/owners"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	JWTmiddleware   middleware.ConfigJwt
	BuyerController buyers.BuyerController
	OwnerController owners.OwnerController
}

func (c1 *ControllerList) RouteRegister(e *echo.Echo) {
	// jwtAuth := middleware.ConfigJwt(c1.JWTmiddleware)
	e.POST("buyers/login", c1.BuyerController.Login)
	e.POST("buyers/register", c1.BuyerController.Register)

	//owners
	e.POST("owners/login", c1.OwnerController.Login)
	e.POST("owners/register", c1.OwnerController.Register)

	//products

}
