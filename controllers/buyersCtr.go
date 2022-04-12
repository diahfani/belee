package controllers

import (
	"belee/config"
	"belee/handler/encrypt"
	"belee/models"
	"belee/models/buyer"
	"errors"
	"strconv"
	"net/http"
	"belee/middlewares"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)


func RegisterController(c echo.Context) error {
	var buyersReg buyer.BuyersRegist
	c.Bind(&buyersReg)
	if buyersReg.Name == "" || buyersReg.Age == "" || buyersReg.NoHp == "" || buyersReg.Dob == "" ||
		buyersReg.Address == "" || buyersReg.Email == "" || buyersReg.Password == "" {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "one or more field are empty",
			Data:    nil,
		})
	}

	var buyersData buyer.Buyers
	buyersData.Name = buyersReg.Name
	buyersData.Age = buyersReg.Age
	buyersData.NoHp = buyersReg.NoHp
	buyersData.Dob = buyersReg.Dob
	buyersData.Address = buyersReg.Address
	buyersData.Email = buyersReg.Email
	buyersData.Password, _ = encrypt.Hash(buyersReg.Password)
	result := config.DB.Create(&buyersData)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "duplicate email",
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
	buyersLogin := buyer.BuyersLogin{}
	c.Bind(&buyersLogin)

	buyers := buyer.Buyers{}

	result := config.DB.First(&buyers, "email = ?", buyersLogin.Email)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusForbidden, models.BaseResponse{
				Code:    http.StatusForbidden,
				Message: "User not found",
				Data:    nil,
			})
		} else {
			return c.JSON(http.StatusInternalServerError, models.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "There's error in server",
				Data:    nil,
			})
		}
	}

	if !encrypt.CheckPasswordHash(buyersLogin.Password, buyers.Password) {
		return c.JSON(http.StatusForbidden, models.BaseResponse{
			Code:    http.StatusForbidden,
			Message: "password didnt match",
		})
	}
	token, err := middlewares.GenerateTokenBuyersJWT(buyers.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "there's mistake in server",
			Data:    nil,
		})
	}

	buyersJwt := buyer.BuyersResponse{
		Id:      buyers.Id,
		Name:    buyers.Name,
		Age:     buyers.Age,
		NoHp:    buyers.NoHp,
		Dob:     buyers.Dob,
		Address: buyers.Address,
		Email:   buyers.Email,
		Token:   token,
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "Succeed",
		Data:    buyersJwt,
	})
}

func DetailsBuyers(c echo.Context) error {
	var buyers buyer.Buyers
	buyerID, _ := strconv.Atoi(c.Param("buyersId"))

	if err := config.DB.Where("id = ?", buyerID).First(&buyers).Error; err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "record not found",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "success get data",
		Data:    buyers,
	})
}
