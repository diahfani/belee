package routes

import (
	// "belee/middlewares"
	// "belee/middlewares"
	"final_project/belee/controllers"

	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
)

func NewRoutes() *echo.Echo {
	e := echo.New()
	// ev1 := e.Group("api/v1/")

	// e.Pre(middleware.RemoveTrailingSlash())

	//buyers
	buyers := e.Group("api/v1/buyers")
	// buyers.GET("", controllers.GetBuyersController)
	buyers.POST("/login", controllers.LoginController)
	buyers.POST("/register", controllers.RegisterController)
	// buyers.GET("/:buyersId", controllers.DetailsBuyers)

	//owners
	owners := e.Group("api/v1/owners")
	// owners.GET("", controllers.GetOwnersController)
	owners.POST("/login", controllers.OwnersLoginController)
	owners.POST("/register", controllers.OwnersRegisterController)
	// owners.GET("/:buyersId", controllers.DetailsOwners)

	//products
	products := e.Group("api/v1/products")
	products.GET("", controllers.GetProducts)
	products.GET("/:productsId", controllers.DetailsProducts)
	// hanya owners
	products.POST("/add", controllers.CreateProducts)
	products.PUT("/update", controllers.UpdateProducts)
	products.DELETE("/delete/:productsId", controllers.DeleteProducts)

	//warungs
	warung := e.Group("api/v1/warungs")
	warung.GET("", controllers.GetWarung)
	warung.GET("/:warungId", controllers.GetDetailsWarung)
	// hanya owner
	warung.POST("/add", controllers.AddWarung)
	warung.PUT("/update", controllers.UpdateWarung)
	warung.DELETE("/delete/:warungId", controllers.DeleteWarung)

	// transactions
	// transaction := e.Group("api/v1/transactions")
	// transaction.POST("", controllers.AddTransactions)
	// transaction.GET("/:Id", controllers.DetailsTransaction)

	//paymentMethod
	// payment := e.Group("api/v1/payment")
	// payment.POST("", controllers.AddPayment)
	// payment.GET("/:payId", controllers.DetailsPayment)

	//productsType
	typeProducts := e.Group("api/v1/typeProducts")
	typeProducts.POST("/add", controllers.AddProductType)
	typeProducts.GET("", controllers.GetProductType)
	typeProducts.GET("/:pTypeId", controllers.GetDetailsProductsType)
	// typeProducts.PUT("/:id", controllers.UpdateTypeProducts)
	// typeProducts.DELETE("/:id", controllers.DeleteTypeProducts)

	// search products in warung
	// warung.GET("/warung/:warungId/products/productsId", controllers.GetProducts)
	return e
}
