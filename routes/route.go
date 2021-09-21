package routes

import (
	"Documents/belee/controllers"

	"github.com/labstack/echo/v4"
)

func NewRoutes() *echo.Echo {
	e := echo.New()
	// ev1 := e.Group("api/v1/")

	//buyers
	buyers := e.Group("api/v1/buyers")
	buyers.GET("", controllers.GetBuyersController)
	buyers.POST("/login", controllers.LoginController)
	buyers.POST("/register", controllers.RegisterController)
	buyers.GET("/:buyersId", controllers.DetailsBuyers)

	//owners
	owners := e.Group("api/v1/owners")
	owners.GET("", controllers.GetOwnersController)
	owners.POST("/login", controllers.OwnersLoginController)
	owners.POST("/register", controllers.OwnersRegisterController)
	owners.GET("/:buyersId", controllers.DetailsOwners)

	//products
	products := e.Group("api/v1/products")
	products.GET("", controllers.GetProducts)
	products.POST("/add", controllers.CreateProducts)
	products.GET("/:productsId", controllers.DetailsProducts)
	products.PUT("/update", controllers.UpdateProducts)
	products.DELETE("/delete/:productsId", controllers.DeleteProducts)

	//warungs
	warung := e.Group("api/v1/warungs")
	warung.GET("", controllers.GetWarung)
	warung.POST("/add", controllers.AddWarung)
	warung.GET("/:warungId", controllers.GetDetailsWarung)
	warung.PUT("/update", controllers.UpdateWarung)
	warung.DELETE("/delete/:warungId", controllers.DeleteWarung)

	return e
}
