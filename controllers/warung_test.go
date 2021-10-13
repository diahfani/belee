package controllers

import (
	"belee/config"
	"belee/models/warung"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func InsertDataWarung() error {
	warung := warung.Warungs{
		Id:       '1',
		OwnersID: '3',
		Name:     "Warcis",
		Address:  "karawang",
	}
	var err error
	if err = config.DB.Save(&warung).Error; err != nil {
		return err
	}
	return nil
}

func TestAddWarungSuccess(t *testing.T) {
	e := InitEchoTestApi()
	// InsertDataProducts()
	requestBody := strings.NewReader(`{
		"ownersId":"2",
		"name":"toko hamida",
		"address":"jakarta"
		}`)
	req := httptest.NewRequest(http.MethodPost, "http://3.144.166.87:8080/api/v1/warungs", requestBody)

	req.Header.Add("Content-Type", "application/json")
	record := httptest.NewRecorder()
	c := e.NewContext(req, record)

	if assert.NoError(t, AddWarung(c)) {
		response := record.Result()
		assert.Equal(t, 200, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 200, int(responseBody["code"].(float64)))
		assert.Equal(t, "success add warung", responseBody["message"])
	}
}

func TestAddWarungFailedEmptyName(t *testing.T) {
	e := InitEchoTestApi()
	// InsertDataProducts()
	requestBody := strings.NewReader(`{
		"ownersId":"2",
		"name":"",
		"address":"jakarta"
		}`)
	req := httptest.NewRequest(http.MethodPost, "http://3.144.166.87:8080/api/v1/warungs", requestBody)

	req.Header.Add("Content-Type", "application/json")
	record := httptest.NewRecorder()
	c := e.NewContext(req, record)

	if assert.NoError(t, AddWarung(c)) {
		response := record.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "Name empty", responseBody["message"])
	}
}

func TestUpdateWarungFailedEmptyParam(t *testing.T) {
	e := InitEchoTestApi()
	// InsertDataProducts()
	requestBody := strings.NewReader(`{
		"ownersId":"2",
		"name":"",
		"address":"jakarta"
		}`)
	req := httptest.NewRequest(http.MethodPut, "http://3.144.166.87:8080/api/v1/warungs/", requestBody)

	req.Header.Add("Content-Type", "application/json")
	record := httptest.NewRecorder()
	c := e.NewContext(req, record)

	if assert.NoError(t, UpdateWarung(c)) {
		response := record.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "param not valid", responseBody["message"])
	}
}

func TestDeleteWarungFailedEmptyParam(t *testing.T) {
	e := InitEchoTestApi()
	// InsertDataProducts()
	requestBody := strings.NewReader(`{
		"ownersId":"2",
		"name":"",
		"address":"jakarta"
		}`)
	req := httptest.NewRequest(http.MethodDelete, "http://3.144.166.87:8080/api/v1/warungs/", requestBody)

	req.Header.Add("Content-Type", "application/json")
	record := httptest.NewRecorder()
	c := e.NewContext(req, record)

	if assert.NoError(t, UpdateWarung(c)) {
		response := record.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "param not valid", responseBody["message"])
	}
}

func TestGetWarungFailed(t *testing.T) {
	e := InitEchoTestApi()
	InsertDataWarung()

	// requestBody := strings.NewReader(`{
	// 	"warungId":"1"
	// }`)
	request := httptest.NewRequest(http.MethodGet, "http://3.144.166.87:8080/api/v1/warungs", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	c.SetPath("/warungs/:Id")
	c.SetParamNames("Id")
	c.SetParamValues("1")
	if assert.NoError(t, GetDetailsWarung(c)) {
		response := recorder.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "record not found", responseBody["message"])
		// assert.Equal(t, 1, responseBody["data"].(map[string]interface{})["id"])
		// assert.Equal(t, "Diah", responseBody["data"].(map[string]interface{})["name"])
		// assert.Equal(t, "diahaaa@aol.com", responseBody["data"].(map[string]interface{})["email"])

	}
}
