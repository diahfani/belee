package controllers

import (
	"belee/models/buyer"
	"errors"
	"final_project/belee/config"
	"final_project/belee/models"

	// "belee/models/buyers"
	"net/http"

	"belee/middlewares"
	// "belee/models/buyer"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// func GetBuyersController(c echo.Context) error {
// 	buyers := []buyers.Buyers{}

// 	result := config.DB.Find(&buyers)

// 	if result.Error != nil {
// 		if result.Error != gorm.ErrRecordNotFound {
// 			return c.JSON(http.StatusInternalServerError, models.BaseResponse{
// 				Code:    http.StatusInternalServerError,
// 				Message: "Error input data",
// 				Data:    nil,
// 			})
// 		}
// 	}
// 	return c.JSON(http.StatusOK, models.BaseResponse{
// 		Code:    http.StatusOK,
// 		Message: "Succeed",
// 		Data:    buyers,
// 	})
// }

func RegisterController(c echo.Context) error {
	var buyersReg buyer.BuyersRegist
	var emailExist buyer.Buyers
	c.Bind(&buyersReg)
	// emailExist := c.Param("email")
	if emailExist.Email != buyersReg.Email {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "duplicate email",
			Data:    nil,
		})
	}
	//validations
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
	buyersLogin := buyer.BuyersLogin{}
	c.Bind(&buyersLogin)

	buyers := buyer.Buyers{}

	result := config.DB.First(&buyers, "email = ? AND password = ?", buyersLogin.Email, buyersLogin.Password)

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

	token, err := middlewares.CreateToken(buyers.Id)
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
	// buyersJwt := {
	// 	Id: buyers.Id,
	// 	Name: buyers.Name,
	// 	Age: buyers.Age,
	// 	NoHp : buyers.NoHp,
	// 	Dob: buyers.Dob,

	// }

	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "Succeed",
		Data:    buyersJwt,
	})
}

// func DetailsBuyers(c echo.Context) error {
// 	buyersId, err := strconv.Atoi(c.Param("buyersId"))
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
// 			Code:    http.StatusInternalServerError,
// 			Message: "failed",
// 			Data:    nil,
// 		})
// 	}
// 	return c.JSON(http.StatusOK, models.BaseResponse{
// 		Code:    http.StatusOK,
// 		Message: "succeed",
// 		Data:    buyers.Buyers{Id: buyersId},
// 	})
// }
