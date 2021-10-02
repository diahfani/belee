package controllers

import (
	"belee/config"
	"belee/models"
	"belee/models/products"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateProducts(c echo.Context) error {
	var addproducts products.AddProducts
	c.Bind(&addproducts)

	if addproducts.BarangName == "" {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Name empty",
			Data:    nil,
		})
	}

	var productsdata products.Products
	// warungdata.OwnersID = addwarungs.OwnersID
	productsdata.BarangName = addproducts.BarangName
	productsdata.Qty = addproducts.Qty
	productsdata.Price = addproducts.Price
	// warungdata.OwnersName = addwarungs.OwnersName
	// var Owners owner.Owners
	// joinOwner := config.DB.Joins("owners").Where("id = ? and name = ?", Owners.Id, Owners.Name)

	// if joinOwner.Error != nil {
	// 	return c.JSON(http.StatusInternalServerError, models.BaseResponse{
	// 		Code:    http.StatusInternalServerError,
	// 		Message: "gabisa join",
	// 		Data:    nil,
	// 	})
	// }
	result := config.DB.Create(&productsdata)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "there's mistake when input data",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "success add products",
		Data:    (&productsdata),
	})
}

func GetProducts(c echo.Context) error {
	products := []products.Products{}
	result := config.DB.Find(&products)

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
		Data:    products,
	})
}

func UpdateProducts(c echo.Context) error {
	var products products.Products
	productsId, err := strconv.Atoi(c.Param("productsId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "param not valid",
			Data:    nil,
		})
	}

	result := config.DB.First(&products, productsId)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "data not found",
			Data:    nil,
		})
	}

	c.Bind(&products)
	result = config.DB.Save(&products)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "cant update data",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "success update data",
		Data:    &products,
	})
}

func DeleteProducts(c echo.Context) error {
	var products products.Products
	productsId, err := strconv.Atoi(c.Param("productsId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "param not valid",
			Data:    nil,
		})
	}

	result := config.DB.First(&products, productsId)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "data not found",
			Data:    nil,
		})
	}

	c.Bind(&products)
	result = config.DB.Where("id = ?", productsId).Unscoped().Delete(&products)
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
		Data:    &products,
	})

}

func DetailsProducts(c echo.Context) error {
	var products products.Products
	productsname := c.Param("productsName")

	if err := config.DB.Where("barang_name = ?", productsname).First(&products).Error; err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "record not found",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "success get data",
		Data:    products,
	})

}
