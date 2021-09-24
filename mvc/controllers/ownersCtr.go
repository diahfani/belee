package controllers

import (
	"errors"
	"final_project/belee/config"
	"final_project/belee/models"
	"final_project/belee/models/owners"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// func GetOwnersController(c echo.Context) error {
// 	owners := []owners.Owners{}

// 	result := config.DB.Find(&owners)

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
// 		Data:    owners,
// 	})
// }

func OwnersRegisterController(c echo.Context) error {
	var ownersReg owners.OwnersRegist
	c.Bind(&ownersReg)
	// emailExist := c.Param("email")

	//validations
	if ownersReg.Name == "" || ownersReg.Age == "" || ownersReg.NoHp == "" || ownersReg.Dob == "" ||
		ownersReg.Address == "" || ownersReg.Email == "" || ownersReg.Password == "" {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "one or more field are empty",
			Data:    nil,
		})
	}

	var ownersData owners.Owners
	ownersData.Name = ownersReg.Name
	ownersData.Age = ownersReg.Age
	ownersData.NoHp = ownersReg.NoHp
	ownersData.Dob = ownersReg.Dob
	ownersData.Address = ownersReg.Address
	ownersData.Email = ownersReg.Email
	ownersData.Password = ownersReg.Password

	result := config.DB.Create(&ownersData)
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
		Data:    (&ownersData),
	})
}

func OwnersLoginController(c echo.Context) error {
	ownersLogin := owners.OwnersLogin{}
	c.Bind(&ownersLogin)

	owners := owners.Owners{}

	result := config.DB.First(&owners, "email = ? AND password = ?", ownersLogin.Email, ownersLogin.Password)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusForbidden, models.BaseResponse{
				Code:    http.StatusForbidden,
				Message: "owner not found",
				Data:    nil,
			})
		} else {
			return c.JSON(http.StatusInternalServerError, models.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Server error",
				Data:    nil,
			})
		}
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "Succeed",
		Data:    ownersLogin,
	})
}

// func DetailsOwners(c echo.Context) error {
// 	ownersID, err := strconv.Atoi(c.Param("ownersId"))
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
// 		Data:    owners.Owners{Id: ownersID},
// 	})
// }
