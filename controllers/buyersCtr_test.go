package controllers

import (
	"belee/config"
	"belee/handler/encrypt"
	"belee/models/buyer"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEchoTestApi() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

func InsertDataBuyers() error {
	realPassword, _ := encrypt.Hash("diahaufa000")
	buyer := buyer.Buyers{
		Name:     "Diah",
		Password: realPassword,
		Email:    "diahaaa@aol.com",
	}

	var err error
	if err = config.DB.Save(&buyer).Error; err != nil {
		return err
	}
	return nil
}

func TestRegisterBuyerSuccess(t *testing.T) {
	e := InitEchoTestApi()

	requestBody := strings.NewReader(`{
		"name":"haidar ahmad",
		"age":"20",
		"nohp":"089789075746",
		"dob":"2001-01-18",
		"address":"jakarta",
		"email":"haidar@aol.com",
		"password":"initesting"
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://3.144.166.87:8080/api/v1/buyers/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, RegisterController(c)) {
		response := recorder.Result()
		assert.Equal(t, 200, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 200, int(responseBody["code"].(float64)))
		assert.Equal(t, "succeed", responseBody["message"])
	}
}

func TestRegisterBuyerFailedDuplicateEmail(t *testing.T) {
	e := InitEchoTestApi()
	InsertDataBuyers()

	requestBody := strings.NewReader(`{
		"name":"haidar ahmad",
		"age":"20",
		"nohp":"089789075746",
		"dob":"2001-01-18",
		"address":"jakarta",
		"email":"diahaaa@aol.com",
		"password":"diahaufa000"
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://3.144.166.87:8080/api/v1/buyers/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, RegisterController(c)) {
		response := recorder.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "duplicate email", responseBody["message"])
	}
}

func TestRegisterBuyerFailedEmptyName(t *testing.T) {
	e := InitEchoTestApi()

	requestBody := strings.NewReader(`{
		"name":"",
		"age":"20",
		"nohp":"089789075746",
		"dob":"2001-01-18",
		"address":"jakarta",
		"email":"haidar@aol.com",
		"password":"initesting"
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://3.144.166.87:8080/api/v1/buyers/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, RegisterController(c)) {
		response := recorder.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "one or more field are empty", responseBody["message"])
	}
}

func TestRegisterBuyerFailedEmptyAge(t *testing.T) {
	e := InitEchoTestApi()

	requestBody := strings.NewReader(`{
		"name":"haidar ahmad",
		"age":"",
		"nohp":"089789075746",
		"dob":"2001-01-18",
		"address":"jakarta",
		"email":"haidar@aol.com",
		"password":"initesting"
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://3.144.166.87:8080/api/v1/buyers/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, RegisterController(c)) {
		response := recorder.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "one or more field are empty", responseBody["message"])
	}
}

func TestRegisterBuyerFailedEmptyNohp(t *testing.T) {
	e := InitEchoTestApi()

	requestBody := strings.NewReader(`{
		"name":"haidar ahmad",
		"age":"20",
		"nohp":"",
		"dob":"2001-01-18",
		"address":"jakarta",
		"email":"haidar@aol.com",
		"password":"initesting"
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://3.144.166.87:8080/api/v1/buyers/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, RegisterController(c)) {
		response := recorder.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "one or more field are empty", responseBody["message"])
	}
}

func TestRegisterBuyerFailedEmptyDob(t *testing.T) {
	e := InitEchoTestApi()

	requestBody := strings.NewReader(`{
		"name":"haidar ahmad",
		"age":"20",
		"nohp":"089789075746",
		"dob":"",
		"address":"jakarta",
		"email":"haidar@aol.com",
		"password":"initesting"
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://3.144.166.87:8080/api/v1/buyers/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, RegisterController(c)) {
		response := recorder.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "one or more field are empty", responseBody["message"])
	}
}

func TestRegisterBuyerFailedEmptyAddress(t *testing.T) {
	e := InitEchoTestApi()

	requestBody := strings.NewReader(`{
		"name":"haidar ahmad",
		"age":"20",
		"nohp":"089789075746",
		"dob":"2001-01-18",
		"address":"",
		"email":"haidar@aol.com",
		"password":"initesting"
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://3.144.166.87:8080/api/v1/buyers/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, RegisterController(c)) {
		response := recorder.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "one or more field are empty", responseBody["message"])
	}
}

func TestRegisterBuyerFailedEmptyEmail(t *testing.T) {
	e := InitEchoTestApi()

	requestBody := strings.NewReader(`{
		"name":"haidar ahmad",
		"age":"20",
		"nohp":"089789075746",
		"dob":"2001-01-18",
		"address":"jakarta",
		"email":"",
		"password":"initesting"
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://3.144.166.87:8080/api/v1/buyers/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, RegisterController(c)) {
		response := recorder.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "one or more field are empty", responseBody["message"])
	}
}

func TestRegisterBuyerFailedEmptyPassword(t *testing.T) {
	e := InitEchoTestApi()

	requestBody := strings.NewReader(`{
		"name":"haidar ahmad",
		"age":"20",
		"nohp":"089789075746",
		"dob":"2001-01-18",
		"address":"jakarta",
		"email":"haidar@aol.com",
		"password":""
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://3.144.166.87:8080/api/v1/buyers/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, RegisterController(c)) {
		response := recorder.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "one or more field are empty", responseBody["message"])
	}
}

func TestLoginBuyersFailedByEmail(t *testing.T) {
	e := InitEchoTestApi()
	InsertDataBuyers()

	requestBody := strings.NewReader(`{
		"email":"diaharini@aol.com",
		"password":"diahaufa000"
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://3.144.166.87:8080/api/v1/buyers/login", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, LoginController(c)) {
		response := recorder.Result()
		assert.Equal(t, 500, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 500, int(responseBody["code"].(float64)))
		assert.Equal(t, "There's error in server", responseBody["message"])

	}
}

func TestLoginBuyersFailedByPassword(t *testing.T) {
	e := InitEchoTestApi()
	InsertDataBuyers()

	requestBody := strings.NewReader(`{
		"email":"diaharini@aol.com",
		"password":"diahaufa012"
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://3.144.166.87:8080/api/v1/buyers/login", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, LoginController(c)) {
		response := recorder.Result()
		assert.Equal(t, 403, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 403, int(responseBody["code"].(float64)))
		assert.Equal(t, "User not found", responseBody["message"])

	}
}

func TestLoginBuyersSuccess(t *testing.T) {
	e := InitEchoTestApi()
	InsertDataBuyers()

	requestBody := strings.NewReader(`{
		"email":"diahaaa@aol.com",
		"password":"diahaufa000"
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://3.144.166.87:8080/api/v1/buyers/login", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, LoginController(c)) {
		response := recorder.Result()
		assert.Equal(t, 200, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 200, int(responseBody["code"].(float64)))
		assert.Equal(t, "Succeed", responseBody["message"])
		// assert.Equal(t, 1, responseBody["data"].(map[string]interface{})["id"])
		assert.Equal(t, "Diah", responseBody["data"].(map[string]interface{})["name"])
		assert.Equal(t, "diahaaa@aol.com", responseBody["data"].(map[string]interface{})["email"])

	}
}
