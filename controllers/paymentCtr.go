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
	// var nameExist paymentMethod.PaymentMethods
	c.Bind(&addPayment)

	// if nameExist.Name == addPayment.Name {
	// 	return c.JSON(http.StatusBadRequest, models.BaseResponse{
	// 		Code:    http.StatusBadRequest,
	// 		Message: "duplicate name/method has been added before",
	// 		Data:    nil,
	// 	})
	// }

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
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "duplicate payment name",
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
				Message: "cant get data",
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
