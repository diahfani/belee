package controllers

import (
	"belee/config"
	"belee/middlewares"
	"belee/models"
	"belee/models/owner"
	"errors"

	"net/http"
	"strconv"

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
	var ownersReg owner.OwnersRegist
	// var emailExist owner.Owners
	c.Bind(&ownersReg)
	// emailExist := c.Param("email")
	// if emailExist.Email == ownersReg.Email {
	// 	return c.JSON(http.StatusBadRequest, models.BaseResponse{
	// 		Code:    http.StatusBadRequest,
	// 		Message: "duplicate email",
	// 		Data:    nil,
	// 	})
	// }

	//validations
	if ownersReg.Name == "" || ownersReg.Age == "" || ownersReg.NoHp == "" || ownersReg.Dob == "" ||
		ownersReg.Address == "" || ownersReg.Email == "" || ownersReg.Password == "" {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "one or more field are empty",
			Data:    nil,
		})
	}

	var ownersData owner.Owners
	ownersData.Name = ownersReg.Name
	ownersData.Age = ownersReg.Age
	ownersData.NoHp = ownersReg.NoHp
	ownersData.Dob = ownersReg.Dob
	ownersData.Address = ownersReg.Address
	ownersData.Email = ownersReg.Email
	ownersData.Password = ownersReg.Password

	// err := ownersData.BeforeSave(config.DB)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	result := config.DB.Create(&ownersData)
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
		Data:    (&ownersData),
	})
}

func OwnersLoginController(c echo.Context) error {
	ownersLogin := owner.OwnersLogin{}
	c.Bind(&ownersLogin)

	owners := owner.Owners{}

	result := config.DB.First(&owners, "email = ?", ownersLogin.Email)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusForbidden, models.BaseResponse{
				Code:    http.StatusForbidden,
				Message: "Owner not found",
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

	// if !encrypt.CheckPasswordHash(ownersLogin.Password, owners.Password) {
	// 	return c.JSON(http.StatusForbidden, models.BaseResponse{
	// 		Code:    http.StatusForbidden,
	// 		Message: "password didnt match",
	// 	})
	// }

	token, err := middlewares.GenerateTokenOwnersJWT(owners.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "there's mistake in server",
			Data:    nil,
		})
	}
	ownersJwt := owner.OwnersResponse{
		Id:      owners.Id,
		Name:    owners.Name,
		Age:     owners.Age,
		NoHp:    owners.NoHp,
		Dob:     owners.Dob,
		Address: owners.Address,
		Email:   owners.Email,
		Token:   token,
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "Succeed",
		Data:    ownersJwt,
	})

}

func DetailsOwners(c echo.Context) error {
	var owners owner.Owners
	ownerID, _ := strconv.Atoi(c.Param("ownersId"))

	if err := config.DB.Where("id = ?", ownerID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "record not found",
			Data:    nil,
		})
	}
	// config.DB.Preload("Warung").Find(&owners)
	// var warung warung.Warungs
	// config.DB.Where("id = ?", owners.Id).Preload("Warungs").First(&warung)
	return c.JSON(http.StatusOK, models.BaseResponse{
		Code:    http.StatusOK,
		Message: "success get data",
		Data:    owners,
	})

	// .Preload("Warungs").First(&owners)
}
