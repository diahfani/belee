package controllers

import (
	"belee/config"
	"belee/models"
	"belee/models/transactions"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddTransaction(c echo.Context) error {
	var addTr transactions.Addtransactions
	c.Bind(&addTr)
	// if err := c.Bind(&addTr); err != nil {
	// 	return c.JSON(http.StatusBadRequest, models.BaseResponse{
	// 		Code:    http.StatusBadRequest,
	// 		Message: "cant create transactions",
	// 		Data:    nil,
	// 	})
	// }
	if addTr.ProductsName == "" {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Name empty",
			Data:    nil,
		})
	}
	var transData transactions.Transactions
	transData.ProductsName = addTr.ProductsName
	transData.TotalQty = addTr.TotalQty
	transData.TotalPrice = addTr.Totalprice
	// transData.Status = addTr.Status

	result := config.DB.Create(&addTr)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "cant add data",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "success create transactions",
		Data:    result,
	})

}

func DetailsTransaction(c echo.Context) error {
	var tr transactions.Transactions
	transId := c.Param("transactionId")

	if err := config.DB.Where("id = ?", transId).Error; err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "record not found",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "success get data",
		Data:    tr,
	})
	// var tr transactions.Transactions
	// trID, _ := strconv.Atoi(c.Param("transactionsId"))

	// if err := config.DB.Where("id = ?", trID).Preload("Buyers", "Warungs", "Products", "PaymentMethods").Find(&tr); err != nil {
	// 	return c.JSON(http.StatusBadRequest, models.BaseResponse{
	// 		Code:    http.StatusBadRequest,
	// 		Message: "record not found",
	// 		Data:    nil,
	// 	})
	// }

	// return c.JSON(http.StatusOK, models.BaseResponse{
	// 	Code:    http.StatusOK,
	// 	Message: "success get data",
	// 	Data:    tr,
	// })
}

func DeleteTransaction(c echo.Context) error {
	var tr transactions.Transactions
	trId, err := strconv.Atoi(c.Param("transactionId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "param not valid",
			Data:    nil,
		})
	}

	result := config.DB.First(&tr, trId)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "data not found",
			Data:    nil,
		})
	}

	c.Bind(&tr)
	result = config.DB.Where("id = ?", trId).Unscoped().Delete(&tr)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "cant delete data",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "success delete data",
		Data:    &tr,
	})

}
