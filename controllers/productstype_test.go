package controllers

import (
	"belee/config"
	"belee/models/productsType"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func InsertProductTypes() error {
	payment := productsType.ProductsType{
		NameType: "makanan",
	}
	var err error
	if err = config.DB.Save(&payment).Error; err != nil {
		return err
	}
	return nil
}

func TestAddProductsTypeSuccess(t *testing.T) {
	e := InitEchoTestApi()
	requestBody := strings.NewReader(`{"nametype":"makanan"}`)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/typeProducts", requestBody)
	// requestBody := strings.NewReader(`{"name":""}`)
	req.Header.Add("Content-Type", "application/json")
	record := httptest.NewRecorder()
	c := e.NewContext(req, record)

	if assert.NoError(t, AddProductType(c)) {
		response := record.Result()
		assert.Equal(t, 200, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 200, int(responseBody["code"].(float64)))
		assert.Equal(t, "success add products type", responseBody["message"])
	}
}

func TestAddProductsTypeFailedEmptyName(t *testing.T) {
	e := InitEchoTestApi()
	requestBody := strings.NewReader(`{"nametype":""}`)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/typeProducts", requestBody)
	// requestBody := strings.NewReader(`{"name":""}`)
	req.Header.Add("Content-Type", "application/json")
	record := httptest.NewRecorder()
	c := e.NewContext(req, record)

	if assert.NoError(t, AddProductType(c)) {
		response := record.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "Name empty", responseBody["message"])
	}
}

func TestGetProductsTypeFailed(t *testing.T) {
	e := InitEchoTestApi()
	InsertProductTypes()
	requestBody := strings.NewReader(`{"id":"1"}`)

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8000/api/v1/typeProducts/", requestBody)
	// requestBody := strings.NewReader(`{"name":""}`)
	req.Header.Add("Content-Type", "application/json")
	record := httptest.NewRecorder()
	c := e.NewContext(req, record)

	if assert.NoError(t, GetDetailsProductsType(c)) {
		response := record.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "record not found", responseBody["message"])
	}
}
