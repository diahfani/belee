// controllers sebagai gerbang
package buyers

import (
	// "final_project/belee/belee/business/buyers"
	// "final_project/belee/belee/controllers"

	"final_project/belee/business/buyers"
	"final_project/belee/controllers"
	_baseresponse "final_project/belee/controllers"
	"final_project/belee/controllers/buyers/request"
	"final_project/belee/controllers/buyers/responses"

	// "final_project/belee/drivers/databases/buyers"

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
	ctx := c.Request().Context()

	req := request.BuyersLogin{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, err)
	}

	buyer, token, err := ctr.BuyerUsecase.Login(ctx, *req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, err)
	}

	response := responses.GetAuthResp(buyer, token)
	return controllers.NewSuccessResponse(c, response)

	// fmt.Println("Login")
	// buyersLogin := request.BuyersLogin{}
	// c.Bind(&buyersLogin)

	// ctx := c.Request().Context()
	// buyer, err := ctr.BuyerUsecase.Login(ctx, buyersLogin.Email, buyersLogin.Password)

	// if err != nil {
	// 	controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	// 	// NewErrorResponse(c, http.StatusInternalServerError, error)
	// }

	// return controllers.NewSuccessResponse(c, responses.FromDomain(buyer))
	// // return .NewSuccessResponse(c, responses.FromDomain(buyer))
}

func (ctr BuyerController) Register(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.BuyersRegist{}
	// req := buyers.Buyers{}
	if err := c.Bind(&req); err != nil {
		return _baseresponse.NewErrorResponse(c, err)
	}

	// reqDomain := req.ToDomain()
	buyer, token, err := ctr.BuyerUsecase.Register(ctx, *req.ToDomain())
	if err != nil {
		return _baseresponse.NewErrorResponse(c, err)
	}
	response := responses.GetAuthResp(buyer, token)
	return _baseresponse.NewSuccessResponse(c, response)
}
