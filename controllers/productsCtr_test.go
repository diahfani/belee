package controllers

import (
	"belee/config"
	"belee/models/products"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// type (
// 	Products struct {
// 		Id           int     `json:"id"`
// 		WarungID     int     `json:"warungid"`
// 		BarangTypeID int     `json:"barangtypeid"`
// 		BarangName   string  `json:"productname"`
// 		Qty          int     `json:"qty"`
// 		Price        float64 `json:"price"`
// 	}
// 	handler struct {
// 		db map[string]*Products
// 	}
// )

// func (h *handler) getProducts(c echo.Context) error {
// 	id := c.Param("id")
// 	product := h.db[id]
// 	if product == nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "param not valid")
// 	}
// 	return c.JSON(http.StatusOK, product)
// }

func InsertDataProducts() error {
	product := products.Products{
		Id:           1,
		WarungID:     16,
		BarangTypeID: 2,
		BarangName:   "lada bubuk",
		Qty:          30,
		Price:        1500,
	}
	var err error
	if err = config.DB.Save(&product).Error; err != nil {
		return err
	}
	return nil
}

func TestAddProductsSuccess(t *testing.T) {
	e := InitEchoTestApi()
	// InsertDataProducts()
	requestBody := strings.NewReader(`{
		"warungId":"16",
		"barangTypeId":"2",
		"name":"lada bubuk",
		"qty":"30",
		"price":"1500"
		}`)
	req := httptest.NewRequest(http.MethodPost, "http://3.144.166.87:8080/api/v1/products", requestBody)

	req.Header.Add("Content-Type", "application/json")
	record := httptest.NewRecorder()
	c := e.NewContext(req, record)

	if assert.NoError(t, CreateProducts(c)) {
		response := record.Result()
		assert.Equal(t, 200, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 200, int(responseBody["code"].(float64)))
		assert.Equal(t, "success add products", responseBody["message"])
	}
}

func TestAddProductsFailedEmptyName(t *testing.T) {
	e := InitEchoTestApi()
	// InsertDataProducts()
	requestBody := strings.NewReader(`{
		"warungId":"16",
		"barangTypeId":"2",
		"name":"",
		"qty":"30",
		"price":"1500"
		}`)
	req := httptest.NewRequest(http.MethodGet, "http://3.144.166.87:8080/api/v1/products", requestBody)

	req.Header.Add("Content-Type", "application/json")
	record := httptest.NewRecorder()
	c := e.NewContext(req, record)

	if assert.NoError(t, CreateProducts(c)) {
		response := record.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "Name empty", responseBody["message"])
	}
}

func TestUpdateProductsFailedParam(t *testing.T) {
	e := InitEchoTestApi()

	// InsertDataProducts()

	// app := echo.New()
	requestBody := strings.NewReader(`{
		"warungId":"16",
		"barangTypeId":"2",
		"name":"garam",
		"qty":"30",
		"price":"1500"
		}`)
	req := httptest.NewRequest(http.MethodPut, "http://3.144.166.87:8080/api/v1/products/update/", requestBody)

	req.Header.Add("Content-Type", "application/json")
	record := httptest.NewRecorder()
	c := e.NewContext(req, record)
	// c.SetPath("update/:id")
	// c.SetParamNames("id")
	// c.SetParamValues("1")
	// h := &handler{}

	if assert.NoError(t, UpdateProducts(c)) {
		response := record.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "param not valid", responseBody["message"])
	}
}

func TestGetProductsSuccess(t *testing.T) {
	e := InitEchoTestApi()

	InsertDataProducts()

	// app := echo.New()
	requestBody := strings.NewReader(`{
		"name":"lada bubuk"`)
	req := httptest.NewRequest(http.MethodGet, "http://3.144.166.87:8080/api/v1/products/", requestBody)

	req.Header.Add("Content-Type", "application/json")
	record := httptest.NewRecorder()
	c := e.NewContext(req, record)
	// c.SetPath("update/:id")
	// c.SetParamNames("id")
	// c.SetParamValues("1")
	// h := &handler{}

	if assert.NoError(t, DetailsProducts(c)) {
		response := record.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "record not found", responseBody["message"])
	}
}

func TestGetProductsFailedParam(t *testing.T) {
	e := InitEchoTestApi()

	InsertDataProducts()

	// app := echo.New()
	requestBody := strings.NewReader(`{
		"name":"garam"`)
	req := httptest.NewRequest(http.MethodGet, "http://3.144.166.87:8080/api/v1/products/", requestBody)

	req.Header.Add("Content-Type", "application/json")
	record := httptest.NewRecorder()
	c := e.NewContext(req, record)
	// c.SetPath("update/:id")
	// c.SetParamNames("id")
	// c.SetParamValues("1")
	// h := &handler{}

	if assert.NoError(t, DetailsProducts(c)) {
		response := record.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "record not found", responseBody["message"])
	}
}

// func TestDeleteProductsSuccess(t *testing.T) {
// 	e := InitEchoTestApi()

// 	InsertDataProducts()

// 	// app := echo.New()
// 	// requestBody := strings.NewReader(`{
// 	// 	"warungId":"16",
// 	// 	"barangTypeId":"2",
// 	// 	"name":"garam",
// 	// 	"qty":"30",
// 	// 	"price":"1500"
// 	// 	}`)
// 	req := httptest.NewRequest(http.MethodDelete, "http://localhost:8000/api/v1/products/delete/1", nil)
// 	req.Header.Add("Content-Type", "application/json")
// 	record := httptest.NewRecorder()
// 	c := e.NewContext(req, record)

// 	if assert.NoError(t, DeleteProducts(c)) {
// 		response := record.Result()
// 		assert.Equal(t, 200, response.StatusCode)
// 		body, _ := io.ReadAll(response.Body)
// 		var responseBody map[string]interface{}
// 		json.Unmarshal(body, &responseBody)

// 		assert.Equal(t, 200, int(responseBody["code"].(float64)))
// 		assert.Equal(t, "success delete data", responseBody["message"])
// 	}
// }

func TestDeleteProductsFailedParam(t *testing.T) {
	e := InitEchoTestApi()

	InsertDataProducts()

	// app := echo.New()
	requestBody := strings.NewReader(`{
		"warungId":"16",
		"barangTypeId":"2",
		"name":"garam",
		"qty":"30",
		"price":"1500"
		}`)
	req := httptest.NewRequest(http.MethodPut, "http://3.144.166.87:8080/api/v1/products/delete/", requestBody)

	req.Header.Add("Content-Type", "application/json")
	record := httptest.NewRecorder()
	c := e.NewContext(req, record)

	if assert.NoError(t, DeleteProducts(c)) {
		response := record.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "param not valid", responseBody["message"])
	}
}
