// controllers sebagai gerbang
package buyers

import (
	// "final_project/belee/belee/business/buyers"
	// "final_project/belee/belee/controllers"

	"final_project/belee/business/buyers"
	"final_project/belee/controllers"
	"final_project/belee/controllers/buyers/request"
	"final_project/belee/controllers/buyers/responses"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BuyerController struct {
	BuyerUsecase buyers.Usecase
}

func NewBuyerController(buyerUsecase buyers.Usecase) *BuyerController {
	return &BuyerController{
		BuyerUsecase: buyerUsecase,
	}
}

func (ctr BuyerController) Login(c echo.Context) error {
	fmt.Println("Login")
	buyersLogin := request.BuyersLogin{}
	c.Bind(&buyersLogin)

	ctx := c.Request().Context()
	buyer, err := ctr.BuyerUsecase.Login(buyersLogin.Email, buyersLogin.Password, ctx)

	if err != nil {
		controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
		// NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(buyer))
	// return .NewSuccessResponse(c, responses.FromDomain(buyer))
}
