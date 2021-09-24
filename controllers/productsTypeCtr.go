package controllers

import (
	"belee/models"
	"belee/models/productsType"
	"final_project/belee/config"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddProductType(c echo.Context) error {
	var addProductType productsType.AddType
	c.Bind(&addProductType)

	if addProductType.NameType == "" {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Name empty",
			Data:    nil,
		})
	}

	var productsType productsType.ProductsType
	productsType.NameType = addProductType.NameType

	result := config.DB.Create(&productsType)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "there's mistake when input data",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "success add warung",
		Data:    (&productsType),
	})
}
