package routes

import (
	"belee/controllers"

	"github.com/labstack/echo/v4"
)

func NewRoutes() *echo.Echo {
	e := echo.New()
	// jwtBuyer := middleware.JWT([]byte(constant.SECRET_JWT_BUYERS))
	// jwtOwner := middleware.JWT([]byte(constant.SECRET_JWT_OWNERS))

	//buyers
	buyers := e.Group("api/v1/buyers")
	buyers.POST("/login", controllers.LoginController)
	buyers.POST("/register", controllers.RegisterController)
	buyers.GET("/:buyersId", controllers.DetailsBuyers)

	//owners
	owners := e.Group("api/v1/owners")
	owners.POST("/login", controllers.OwnersLoginController)
	owners.POST("/register", controllers.OwnersRegisterController)
	owners.GET("/:ownersId", controllers.DetailsOwners)

	//products

	products := e.Group("api/v1/products")
	products.GET("", controllers.GetProducts)
	products.GET("/:productsName", controllers.DetailsProducts)
	// hanya owners
	products.POST("/add", controllers.CreateProducts)
	products.PUT("/update/:productsId", controllers.UpdateProducts)
	products.DELETE("/delete/:productsId", controllers.DeleteProducts)

	//warungs
	warung := e.Group("api/v1/warungs")
	warung.GET("", controllers.GetWarung)
	warung.GET("/:warungId", controllers.GetDetailsWarung)
	// hanya owner
	warung.POST("/add", controllers.AddWarung)
	warung.PUT("/update/:warungId", controllers.UpdateWarung)
	warung.DELETE("/delete/:warungId", controllers.DeleteWarung)

	// transactions
	transaction := e.Group("api/v1/transactions")
	transaction.POST("/add", controllers.AddTransaction)
	transaction.GET("/:transactionId", controllers.DetailsTransaction)
	transaction.DELETE("delete/:transactionId", controllers.DeleteTransaction)

	//paymentMethod
	payment := e.Group("api/v1/payment")
	payment.POST("/add", controllers.AddPayment)
	payment.GET("", controllers.GetPayment)

	//productsType
	typeProducts := e.Group("api/v1/typeProducts")
	typeProducts.POST("/add", controllers.AddProductType)
	typeProducts.GET("", controllers.GetProductType)
	typeProducts.GET("/:pTypeId", controllers.GetDetailsProductsType)

	// m.LogMiddleware(e)
	return e
}
