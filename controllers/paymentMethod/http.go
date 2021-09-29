package paymentmethod

import (
	"belee/business/paymentMethod"
	"belee/controllers"
	"belee/controllers/paymentMethod/request"
	"belee/controllers/paymentMethod/responses"

	"github.com/labstack/echo/v4"
)

type PaymentController struct {
	PaymentUsecase paymentMethod.Usecase
}

func NewPaymentController(paymentUc paymentMethod.Usecase) *PaymentController {
	return &PaymentController{
		PaymentUsecase: paymentUc,
	}
}

func (ctr *PaymentController) Add(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.PaymentMethod{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, err)
	}

	resp, err := ctr.PaymentUsecase.Add(ctx, req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(resp))

}

func (ctr *PaymentController) FindAll(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := ctr.PaymentUsecase.FindAll(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, err)
	}

	responseCtr := []responses.PaymentMethod{}
	for _, value := range resp {
		responseCtr = append(responseCtr, responses.FromDomain(value))
	}

	return controllers.NewSuccessResponse(c, responseCtr)
}
