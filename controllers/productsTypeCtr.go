package controllers

import (
	"belee/config"
	"belee/models"
	"belee/models/productsType"
	"net/http"
	"strconv"

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
		Message: "success add products type",
		Data:    (&productsType),
	})
}

func GetDetailsProductsType(c echo.Context) error {
	var productType productsType.ProductsType
	pTypeId, _ := strconv.Atoi(c.Param("pTypeId"))

	if err := config.DB.Where("id = ?", pTypeId).First(&productType).Error; err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "record not found",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "success get data",
		Data:    productType,
	})

}

func GetProductType(c echo.Context) error {
	pType := []productsType.ProductsType{}
	result := config.DB.Find(&pType)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "cant get data",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "success get data",
		Data:    pType,
	})
}
