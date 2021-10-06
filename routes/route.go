package routes

import (
	"belee/constant"
	"belee/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRoutes() *echo.Echo {
	e := echo.New()

	jwtBuyer := middleware.JWT([]byte(constant.SECRET_JWT_BUYERS))
	jwtOwner := middleware.JWT([]byte(constant.SECRET_JWT_OWNERS))

	//buyers
	buyers := e.Group("api/v1/buyers")
	buyers.POST("/login", controllers.LoginController)
	buyers.POST("/register", controllers.RegisterController)
	buyers.GET("/:buyersId", controllers.DetailsBuyers, jwtBuyer)

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
	products.POST("", controllers.CreateProducts, jwtOwner)
	products.PUT("/:productsId", controllers.UpdateProducts, jwtOwner)
	products.DELETE("/:productsId", controllers.DeleteProducts, jwtOwner)

	//warungs
	warung := e.Group("api/v1/warungs")
	warung.GET("", controllers.GetWarung)
	warung.GET("/:warungId", controllers.GetDetailsWarung)
	// hanya owner
	warung.POST("", controllers.AddWarung, jwtOwner)
	warung.PUT("/:warungId", controllers.UpdateWarung, jwtOwner)
	warung.DELETE("/delete/:warungId", controllers.DeleteWarung, jwtOwner)

	// transactions
	transaction := e.Group("api/v1/transactions")
	transaction.POST("", controllers.AddTransaction)
	transaction.GET("/:transactionId", controllers.DetailsTransaction)
	transaction.DELETE("/:transactionId", controllers.DeleteTransaction)

	//paymentMethod
	payment := e.Group("api/v1/payment")
	payment.POST("", controllers.AddPayment)
	payment.GET("", controllers.GetPayment)

	//productsType
	typeProducts := e.Group("api/v1/typeProducts")
	typeProducts.POST("", controllers.AddProductType)
	typeProducts.GET("", controllers.GetProductType)
	typeProducts.GET("/:pTypeId", controllers.GetDetailsProductsType)

	// m.LogMiddleware(e)
	return e
}
