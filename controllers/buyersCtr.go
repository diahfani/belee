package controllers

import (
	"Documents/belee/config"
	"Documents/belee/models"
	"Documents/belee/models/buyers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetBuyersController(c echo.Context) error {
	buyers := []buyers.Buyers{}

	result := config.DB.Find(&buyers)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, models.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error input data",
				Data:    nil,
			})
		}
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "Succeed",
		Data:    buyers,
	})
}

func RegisterController(c echo.Context) error {
	var buyersReg buyers.BuyersRegist
	c.Bind(&buyersReg)

	//validations
	if buyersReg.Name == "" || buyersReg.Age == "" || buyersReg.NoHp == "" || buyersReg.Dob == "" ||
		buyersReg.Address == "" || buyersReg.Email == "" || buyersReg.Password == "" {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "one or more field are empty",
			Data:    nil,
		})
	}

	var buyersData buyers.Buyers
	buyersData.Name = buyersReg.Name
	buyersData.Age = buyersReg.Age
	buyersData.NoHp = buyersReg.NoHp
	buyersData.Dob = buyersReg.Dob
	buyersData.Address = buyersReg.Address
	buyersData.Email = buyersReg.Email
	buyersData.Password = buyersReg.Password

	result := config.DB.Create(&buyersData)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "there's mistake when input data",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "succeed",
		Data:    (&buyersData),
	})
}

func LoginController(c echo.Context) error {
	buyersLogin := buyers.BuyersLogin{}
	c.Bind(&buyersLogin)

	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "Succeed",
		Data:    buyersLogin,
	})
}

func DetailsBuyers(c echo.Context) error {
	buyersId, err := strconv.Atoi(c.Param("buyersId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "failed",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "succeed",
		Data:    buyers.Buyers{Id: buyersId},
	})
}
