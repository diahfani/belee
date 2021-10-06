package controllers

import (
	"belee/config"
	"belee/models/transactions"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func InsertDataTransactions() error {
	transaction := transactions.Transactions{
		Id:           1,
		BuyerID:      1,
		WarungID:     1,
		BarangID:     1,
		PaymentID:    1,
		ProductsName: "lada bubuk",
		TotalQty:     2,
		TotalPrice:   3000,
	}
	var err error
	if err = config.DB.Save(&transaction).Error; err != nil {
		return err
	}
	return nil
}

func TestAddTransactionSuccess(t *testing.T) {
	e := InitEchoTestApi()
	// InsertDataProducts()
	requestBody := strings.NewReader(`{
		"buyerId":"2",
		"warungId":"16",
		"barangId":"2",
		"paymentId":"1",
		"productsname":"lada bubuk",
		"totalqty":"2",
		"totalprice":"3000"
		}`)
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/transactions", requestBody)

	req.Header.Add("Content-Type", "application/json")
	record := httptest.NewRecorder()
	c := e.NewContext(req, record)

	if assert.NoError(t, AddTransaction(c)) {
		response := record.Result()
		assert.Equal(t, 200, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 200, int(responseBody["code"].(float64)))
		assert.Equal(t, "success add transactions", responseBody["message"])
	}
}

func TestAddTransactionFailedEmptyName(t *testing.T) {
	e := InitEchoTestApi()
	// InsertDataProducts()
	requestBody := strings.NewReader(`{
		"buyerId":"2",
		"warungId":"16",
		"barangId":"2",
		"paymentId":"1",
		"productsname":"",
		"totalqty":"2",
		"totalprice":"3000"
		}`)
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/transactions", requestBody)

	req.Header.Add("Content-Type", "application/json")
	record := httptest.NewRecorder()
	c := e.NewContext(req, record)

	if assert.NoError(t, AddTransaction(c)) {
		response := record.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "choose products first", responseBody["message"])
	}
}

func TestGetTransactionFailedParamError(t *testing.T) {
	e := InitEchoTestApi()
	InsertDataProducts()
	requestBody := strings.NewReader(`{
		"id":"1"}`)
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8000/api/v1/transactions/", requestBody)

	req.Header.Add("Content-Type", "application/json")
	record := httptest.NewRecorder()
	c := e.NewContext(req, record)

	if assert.NoError(t, DetailsTransaction(c)) {
		response := record.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "record not found", responseBody["message"])
	}
}

func TestDeleteTransactionFailedParamError(t *testing.T) {
	e := InitEchoTestApi()
	InsertDataProducts()
	requestBody := strings.NewReader(`{
		"id":"1"}`)
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8000/api/v1/transactions/", requestBody)

	req.Header.Add("Content-Type", "application/json")
	record := httptest.NewRecorder()
	c := e.NewContext(req, record)

	if assert.NoError(t, DeleteTransaction(c)) {
		response := record.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "param not valid", responseBody["message"])
	}
}
