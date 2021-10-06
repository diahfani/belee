package controllers

import (
	"belee/config"
	"belee/models"
	"belee/models/products"
	"belee/models/productsType"
	"belee/models/warung"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateProducts(c echo.Context) error {
	var addproducts products.AddProducts
	var warung warung.Warungs
	var producttype productsType.ProductsType
	c.Bind(&addproducts)

	if addproducts.BarangName == "" {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Name empty",
			Data:    nil,
		})
	}

	var productsdata products.Products
	productsdata.BarangTypeID = addproducts.BarangTypeID
	productsdata.WarungID = addproducts.WarungID
	productsdata.BarangName = addproducts.BarangName
	productsdata.Qty = addproducts.Qty
	productsdata.Price = addproducts.Price

	result := config.DB.Create(&productsdata)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "there's mistake when input data",
			Data:    nil,
		})
	}
	config.DB.Preload("warungs").First(&warung, "id = ?", productsdata.WarungID)
	config.DB.Preload("products").First(&producttype, "id = ?", productsdata.BarangTypeID)

	response := products.ProductResponse{
		Id: productsdata.Id,
		// Warung:     &warung,
		BarangType: &producttype,
		BarangName: productsdata.BarangName,
		Qty:        productsdata.Qty,
		Price:      productsdata.Price,
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "success add products",
		Data:    (&response),
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
	var product products.Products
	// var warung warung.Warungs
	// var producttype productsType.ProductsType
	productsname := c.Param("productsName")

	if err := config.DB.Where("barang_name = ?", productsname).First(&product).Error; err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "record not found",
			Data:    nil,
		})
	}
	// config.DB.Preload("buyers").First(&warung, "id = ?", product.WarungID)
	// config.DB.Preload("products").First(&producttype, "id = ?", product.BarangTypeID)

	// response := products.ProductResponse{
	// 	Id:         product.Id,
	// 	Warung:     &warung,
	// 	BarangType: &producttype,
	// 	BarangName: product.BarangName,
	// 	Qty:        product.Qty,
	// 	Price:      product.Price,

	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "success get data",
		Data:    product,
	})

}
