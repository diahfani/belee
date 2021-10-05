package controllers

import (
	"belee/config"
	"belee/models"
	"belee/models/buyer"
	"belee/models/paymentMethod"
	"belee/models/products"
	"belee/models/transactions"
	"belee/models/warung"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddTransaction(c echo.Context) error {
	var addTransaction transactions.Addtransactions

	var buyer buyer.Buyers
	var barang products.Products
	var warung warung.Warungs
	var payment paymentMethod.PaymentMethods
	c.Bind(&addTransaction)

	// prod := config.DB.Where("id = ?", addTransaction.BarangId).Find(&barang).Error
	// if prod != nil {
	// 	return c.JSON(http.StatusBadRequest, models.BaseResponse{
	// 		Code:    http.StatusBadRequest,
	// 		Message: "products not found",
	// 		Data:    nil,
	// 	})
	// }

	if addTransaction.ProductsName == "" {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "choose products first",
			Data:    nil,
		})
	}
	var transactionsData transactions.Transactions
	transactionsData.BuyerID = addTransaction.BuyerId
	transactionsData.WarungID = addTransaction.WarungId
	transactionsData.PaymentID = addTransaction.PaymentId
	transactionsData.BarangID = addTransaction.BarangId
	transactionsData.ProductsName = addTransaction.ProductsName
	transactionsData.TotalQty = addTransaction.TotalQty
	transactionsData.TotalPrice = addTransaction.Totalprice

	result := config.DB.Create(&transactionsData)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "there's mistake when input data",
			Data:    nil,
		})
	}

	// config.DB.Save(&transactionsData)
	config.DB.Preload("buyers").First(&buyer, "id = ?", transactionsData.BuyerID)
	config.DB.Preload("products").First(&barang, "id = ?", transactionsData.BarangID)
	config.DB.Preload("payment_methods").First(&payment, "id = ?", transactionsData.PaymentID)
	config.DB.Preload("warungs").First(&warung, "id = ?", transactionsData.WarungID)
	// TransResponse := transactions.BuyerTransactionResponse{
	// 	Id:    buyer.Id,
	// 	Name:  buyer.Name,
	// 	Email: buyer.Email,
	// }

	response := transactions.TransactionResponse{
		Id:           transactionsData.Id,
		Buyer:        &buyer,
		Warung:       &warung,
		Barang:       &barang,
		Payment:      &payment,
		ProductsName: addTransaction.ProductsName,
		TotalQty:     addTransaction.TotalQty,
		TotalPrice:   addTransaction.Totalprice,
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "success add transactions",
		Data:    (&response),
	})

}

func DetailsTransaction(c echo.Context) error {
	var transaction transactions.Transactions
	var buyer buyer.Buyers
	var barang products.Products
	var warung warung.Warungs
	var payment paymentMethod.PaymentMethods
	transID := c.Param("transactionId")

	if err := config.DB.Where("id = ?", transID).First(&transaction).Error; err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "record not found",
			Data:    nil,
		})
	}
	config.DB.Preload("buyers").First(&buyer, "id = ?", transaction.BuyerID)
	config.DB.Preload("products").First(&barang, "id = ?", transaction.BarangID)
	config.DB.Preload("payment_methods").First(&payment, "id = ?", transaction.PaymentID)
	config.DB.Preload("warungs").First(&warung, "id = ?", transaction.WarungID)

	response := transactions.TransactionResponse{
		Id:           transaction.Id,
		Buyer:        &buyer,
		Warung:       &warung,
		Barang:       &barang,
		Payment:      &payment,
		ProductsName: transaction.ProductsName,
		TotalQty:     transaction.TotalQty,
		TotalPrice:   transaction.TotalPrice,
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "success get data",
		Data:    response,
	})
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
