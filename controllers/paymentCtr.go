package controllers

import (
	"belee/config"
	"belee/models"
	"belee/models/paymentMethod"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AddPayment(c echo.Context) error {
	var addPayment paymentMethod.AddPayment
	c.Bind(&addPayment)

	if addPayment.Name == "" {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Name empty",
			Data:    nil,
		})
	}

	var paymentmethod paymentMethod.PaymentMethods
	paymentmethod.Name = addPayment.Name

	result := config.DB.Create(&paymentmethod)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "there's mistake when input data",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "success add payment method",
		Data:    (&paymentmethod),
	})

}

func GetPayment(c echo.Context) error {
	payment := []paymentMethod.PaymentMethods{}
	result := config.DB.Find(&payment)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, models.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "can't find data",
				Data:    nil,
			})
		}
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "success get data",
		Data:    payment,
	})
}
