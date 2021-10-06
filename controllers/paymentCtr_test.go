package controllers

import (
	"belee/config"
	"belee/models/paymentMethod"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func InsertDataPayment() error {
	payment := paymentMethod.PaymentMethods{
		Name: "cash",
	}
	var err error
	if err = config.DB.Save(&payment).Error; err != nil {
		return err
	}
	return nil
}

func TestAddPaymentSuccess(t *testing.T) {
	e := InitEchoTestApi()
	requestBody := strings.NewReader(`{"name":"cash"}`)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/payment/add", requestBody)
	// requestBody := strings.NewReader(`{"name":""}`)
	req.Header.Add("Content-Type", "application/json")
	record := httptest.NewRecorder()
	c := e.NewContext(req, record)

	if assert.NoError(t, AddPayment(c)) {
		response := record.Result()
		assert.Equal(t, 200, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 200, int(responseBody["code"].(float64)))
		assert.Equal(t, "success add payment method", responseBody["message"])
	}
}

func TestAddPaymentFailedEmptyName(t *testing.T) {
	e := InitEchoTestApi()
	requestBody := strings.NewReader(`{"name":""}`)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/payment/add", requestBody)
	// requestBody := strings.NewReader(`{"name":""}`)
	req.Header.Add("Content-Type", "application/json")
	record := httptest.NewRecorder()
	c := e.NewContext(req, record)

	if assert.NoError(t, AddPayment(c)) {
		response := record.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "Name empty", responseBody["message"])
	}
}

func TestAddPaymentFailedDuplicateName(t *testing.T) {
	e := InitEchoTestApi()
	InsertDataPayment()
	requestBody := strings.NewReader(`{"name":"cash"}`)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/payment/add", requestBody)
	// requestBody := strings.NewReader(`{"name":""}`)
	req.Header.Add("Content-Type", "application/json")
	record := httptest.NewRecorder()
	c := e.NewContext(req, record)

	if assert.NoError(t, AddPayment(c)) {
		response := record.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "duplicate payment name", responseBody["message"])
	}
}

// func TestGetPaymentSuccess(t *testing.T) {
// 	var testcases = []struct{
// 		name string
// 	}

// 	e := InitEchoTestApi()
// 	InsertDataPayment()
// 	req := httptest.NewRequest(http.MethodGet, "http://localhost:8000/api/v1/payment", nil)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// }
