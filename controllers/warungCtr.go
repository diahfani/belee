package controllers

import (
	"belee/config"
	"belee/models"
	"belee/models/owner"
	"belee/models/warung"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddWarung(c echo.Context) error {
	var addwarungs warung.AddWarungs
	var owners owner.Owners
	c.Bind(&addwarungs)

	if addwarungs.Name == "" {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Name empty",
			Data:    nil,
		})
	}

	var warungdata warung.Warungs
	warungdata.OwnersID = addwarungs.OwnersID
	warungdata.Name = addwarungs.Name
	warungdata.Address = addwarungs.Address
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
	result := config.DB.Create(&warungdata)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "there's mistake when input data",
			Data:    nil,
		})
	}
	config.DB.Preload("owners").First(&owners, "id = ?", warungdata.OwnersID)

	response := warung.WarungRes{
		Owner:     &owners,
		Name:      warungdata.Name,
		Address:   warungdata.Address,
		CreatedAt: warungdata.CreatedAt,
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "success add warung",
		Data:    (&response),
	})
}

func GetDetailsWarung(c echo.Context) error {
	var warungs warung.Warungs
	var owners owner.Owners
	warungID, _ := strconv.Atoi(c.Param("warungId"))

	if err := config.DB.Where("id = ?", warungID).First(&warungs).Error; err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "record not found",
			Data:    nil,
		})
	}

	config.DB.Preload("owners").First(&owners, "id = ?", warungs.OwnersID)

	response := warung.WarungRes{
		Owner:     &owners,
		Name:      warungs.Name,
		Address:   warungs.Address,
		CreatedAt: warungs.CreatedAt,
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "success get data",
		Data:    response,
	})

}

func DeleteWarung(c echo.Context) error {
	var warung warung.Warungs
	warungId, err := strconv.Atoi(c.Param("warungId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "param not valid",
			Data:    nil,
		})
	}

	result := config.DB.First(&warung, warungId)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "data not found",
			Data:    nil,
		})
	}

	c.Bind(&warung)
	result = config.DB.Where("id = ?", warungId).Unscoped().Delete(&warung)
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
		Data:    &warung,
	})

}

func GetWarung(c echo.Context) error {
	warung := []warung.Warungs{}
	result := config.DB.Find(&warung)

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
		Data:    warung,
	})
}

func UpdateWarung(c echo.Context) error {
	var warung warung.Warungs
	warungId, err := strconv.Atoi(c.Param("warungId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "param not valid",
			Data:    nil,
		})
	}

	result := config.DB.First(&warung, warungId)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "data not found",
			Data:    nil,
		})
	}

	c.Bind(&warung)
	result = config.DB.Save(&warung)
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
		Data:    &warung,
	})

}

// return echo.ErrBadGateway
// 	return echo.ErrBadGateway
