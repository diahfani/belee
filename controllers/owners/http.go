// controllers sebagai gerbang
package owners

import (
	// "final_project/belee/belee/business/buyers"
	// "final_project/belee/belee/controllers"

	"final_project/belee/business/owners"
	"final_project/belee/controllers"
	_baseresponse "final_project/belee/controllers"
	"final_project/belee/controllers/owners/request"
	"final_project/belee/controllers/owners/responses"

	// "final_project/belee/drivers/databases/buyers"

	"github.com/labstack/echo/v4"
)

type OwnerController struct {
	OwnerUsecase owners.Usecase
}

func NewOwnerController(ownerUsecase owners.Usecase) *OwnerController {
	return &OwnerController{
		OwnerUsecase: ownerUsecase,
	}
}

func (ctr OwnerController) Login(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.OwnersLogin{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, err)
	}

	owner, token, err := ctr.OwnerUsecase.Login(ctx, *req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, err)
	}

	response := responses.GetAuthResp(owner, token)
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

func (ctr OwnerController) Register(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.OwnersRegist{}
	// req := buyers.Buyers{}
	if err := c.Bind(&req); err != nil {
		return _baseresponse.NewErrorResponse(c, err)
	}

	// reqDomain := req.ToDomain()
	owner, token, err := ctr.OwnerUsecase.Register(ctx, *req.ToDomain())
	if err != nil {
		return _baseresponse.NewErrorResponse(c, err)
	}
	response := responses.GetAuthResp(owner, token)
	return _baseresponse.NewSuccessResponse(c, response)
}
