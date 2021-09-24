package controllers

import (
	"final_project/belee/config"
	"final_project/belee/models"
	"final_project/belee/models/warung"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddWarung(c echo.Context) error {
	var addwarungs warung.AddWarungs
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
	warungdata.OwnersName = addwarungs.OwnersName

	result := config.DB.Create(&warungdata)
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
		Data:    (&warungdata),
	})
}

func GetDetailsWarung(c echo.Context) error {
	return echo.ErrBadGateway
}

func DeleteWarung(c echo.Context) error {
	return echo.ErrBadGateway
}

func GetWarung(c echo.Context) error {
	return echo.ErrBadGateway
}

func UpdateWarung(c echo.Context) error {
	return echo.ErrBadGateway
}
