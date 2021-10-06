package controllers

import (
	"belee/config"
	"belee/handler/encrypt"
	"belee/models/owner"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func InsertDataOwners() error {
	realPassword, _ := encrypt.Hash("fani110")
	owner := owner.Owners{
		Name:     "Fani",
		Password: realPassword,
		Email:    "fani@aol.com",
	}

	var err error
	if err = config.DB.Save(&owner).Error; err != nil {
		return err
	}
	return nil
}

func TestRegisterOwnerSuccess(t *testing.T) {
	e := InitEchoTestApi()

	requestBody := strings.NewReader(`{
		"name":"diah aufa",
		"age":"21",
		"nohp":"085890703579",
		"dob":"2000-10-01",
		"address":"bekasi",
		"email":"diahfani@aol.com",
		"password":"initesting"
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/owners/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, OwnersRegisterController(c)) {
		response := recorder.Result()
		assert.Equal(t, 200, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 200, int(responseBody["code"].(float64)))
		assert.Equal(t, "succeed", responseBody["message"])
	}
}

func TestRegisterOwnerFailedEmptyName(t *testing.T) {
	e := InitEchoTestApi()

	requestBody := strings.NewReader(`{
		"name":"",
		"age":"21",
		"nohp":"085890703579",
		"dob":"2000-10-01",
		"address":"bekasi",
		"email":"diahfani@aol.com",
		"password":"initesting"
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/owners/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, OwnersRegisterController(c)) {
		response := recorder.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "one or more field are empty", responseBody["message"])
	}
}

func TestRegisterOwnerFailedEmptyAge(t *testing.T) {
	e := InitEchoTestApi()

	requestBody := strings.NewReader(`{
		"name":"diah aufa",
		"age":"",
		"nohp":"085890703579",
		"dob":"2000-10-01",
		"address":"bekasi",
		"email":"diahfani@aol.com",
		"password":"initesting"
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/owners/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, OwnersRegisterController(c)) {
		response := recorder.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "one or more field are empty", responseBody["message"])
	}
}

func TestRegisterOwnersFailedEmptyNohp(t *testing.T) {
	e := InitEchoTestApi()

	requestBody := strings.NewReader(`{
		"name":"diah aufa",
		"age":"21",
		"nohp":"",
		"dob":"2000-10-01",
		"address":"bekasi",
		"email":"diahfani@aol.com",
		"password":"initesting"
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/owners/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, OwnersRegisterController(c)) {
		response := recorder.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "one or more field are empty", responseBody["message"])
	}
}

func TestRegisterOwnersFailedEmptyDob(t *testing.T) {
	e := InitEchoTestApi()

	requestBody := strings.NewReader(`{
		"name":"diah aufa",
		"age":"21",
		"nohp":"085890703579",
		"dob":"",
		"address":"bekasi",
		"email":"diahfani@aol.com",
		"password":"initesting"
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/owners/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, OwnersRegisterController(c)) {
		response := recorder.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "one or more field are empty", responseBody["message"])
	}
}

func TestRegisterOwnersFailedEmptyAddress(t *testing.T) {
	e := InitEchoTestApi()

	requestBody := strings.NewReader(`{
		"name":"diah aufa",
		"age":"21",
		"nohp":"085890703579",
		"dob":"2000-10-01",
		"address":"",
		"email":"diahfani@aol.com",
		"password":"initesting"
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/owners/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, OwnersRegisterController(c)) {
		response := recorder.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "one or more field are empty", responseBody["message"])
	}
}

func TestRegisterOwnersFailedEmailDuplicate(t *testing.T) {
	e := InitEchoTestApi()
	InsertDataOwners()

	requestBody := strings.NewReader(`{
		"name":"diah aufa",
		"age":"21",
		"nohp":"085890703579",
		"dob":"2000-10-01",
		"address":"bekasi",
		"email":"fani@aol.com",
		"password":"fani110"
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/owners/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, OwnersRegisterController(c)) {
		response := recorder.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "duplicate email", responseBody["message"])
	}
}

func TestRegisterOwnersFailedEmailEmpty(t *testing.T) {
	e := InitEchoTestApi()

	requestBody := strings.NewReader(`{
		"name":"diah aufa",
		"age":"21",
		"nohp":"085890703579",
		"dob":"2000-10-01",
		"address":"bekasi",
		"email":"",
		"password":"initesting"
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/owners/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, OwnersRegisterController(c)) {
		response := recorder.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "one or more field are empty", responseBody["message"])
	}
}

func TestRegisterOwnersFailedEmptyPassword(t *testing.T) {
	e := InitEchoTestApi()

	requestBody := strings.NewReader(`{
		"name":"diah aufa",
		"age":"21",
		"nohp":"085890703579",
		"dob":"2000-10-01",
		"address":"bekasi",
		"email":"diahfani@aol.com",
		"password":""
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/owners/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)
	if assert.NoError(t, OwnersRegisterController(c)) {
		response := recorder.Result()
		assert.Equal(t, 400, response.StatusCode)
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "one or more field are empty", responseBody["message"])
	}
}

// func TestLoginOwnersFailedByEmail(t *testing.T) {
// 	e := InitEchoTestApi()
// 	InsertDataOwners()

// 	requestBody := strings.NewReader(`{
// 		"email":"aufa@aol.com",
// 		"password":"fani110"
// 	}`)
// 	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/owners/login", requestBody)
// 	request.Header.Add("Content-Type", "application/json")
// 	recorder := httptest.NewRecorder()
// 	c := e.NewContext(request, recorder)
// 	if assert.NoError(t, OwnersLoginController(c)) {
// 		response := recorder.Result()
// 		assert.Equal(t, 403, response.StatusCode)
// 		body, _ := io.ReadAll(response.Body)
// 		var responseBody map[string]interface{}
// 		json.Unmarshal(body, &responseBody)

// 		assert.Equal(t, 403, int(responseBody["code"].(float64)))
// 		assert.Equal(t, "Owner not found", responseBody["message"])

// 	}
// }

// // func TestLoginOwnersFailedByPassword(t *testing.T) {
// // 	e := InitEchoTestApi()
// // 	InsertDataOwners()

// // 	requestBody := strings.NewReader(`{
// // 		"email":"fani@aol.com",
// // 		"password":""
// // 	}`)
// // 	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/api/v1/owners/register", requestBody)
// // 	request.Header.Add("Content-Type", "application/json")
// // 	recorder := httptest.NewRecorder()
// // 	c := e.NewContext(request, recorder)
// // 	if assert.NoError(t, OwnersLoginController(c)) {
// // 		response := recorder.Result()
// // 		assert.Equal(t, 403, response.StatusCode)
// // 		body, _ := io.ReadAll(response.Body)
// // 		var responseBody map[string]interface{}
// // 		json.Unmarshal(body, &responseBody)

// // 		assert.Equal(t, 403, int(responseBody["code"].(float64)))
// // 		assert.Equal(t, "password didnt match", responseBody["message"])

// // 	}
// // }

// // func TestLoginBuyersSuccess(t *testing.T) {
// // 	e := InitEchoTestApi()
// // 	InsertDataOwners()

// // 	requestBody := strings.NewReader(`{
// // 		"email":"fani@aol.com",
// // 		"password":"fani110"
// // 	}`)
// // 	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/owners/login", requestBody)
// // 	request.Header.Add("Content-Type", "application/json")
// // 	recorder := httptest.NewRecorder()
// // 	c := e.NewContext(request, recorder)
// // 	if assert.NoError(t, OwnersLoginController(c)) {
// // 		response := recorder.Result()
// // 		assert.Equal(t, 200, response.StatusCode)
// // 		body, _ := io.ReadAll(response.Body)
// // 		var responseBody map[string]interface{}
// // 		json.Unmarshal(body, &responseBody)

// // 		assert.Equal(t, 200, int(responseBody["code"].(float64)))
// // 		assert.Equal(t, "Succeed", responseBody["message"])
// // 		// assert.Equal(t, 1, responseBody["data"].(map[string]interface{})["id"])
// // 		assert.Equal(t, "Fani", responseBody["data"].(map[string]interface{})["name"])
// // 		assert.Equal(t, "fani@aol.com", responseBody["data"].(map[string]interface{})["email"])

// // 	}
// // }

// func TestLoginOwnersSuccess(t *testing.T) {
// 	e := InitEchoTestApi()
// 	InsertDataOwners()

// 	requestBody := strings.NewReader(`{
// 		"email":"fani@aol.com",
// 		"password":"fani110"
// 	}`)
// 	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/owners/login", requestBody)
// 	request.Header.Add("Content-Type", "application/json")
// 	recorder := httptest.NewRecorder()
// 	c := e.NewContext(request, recorder)
// 	if assert.NoError(t, OwnersLoginController(c)) {
// 		response := recorder.Result()
// 		assert.Equal(t, 200, response.StatusCode)
// 		body, _ := io.ReadAll(response.Body)
// 		var responseBody map[string]interface{}
// 		json.Unmarshal(body, &responseBody)

// 		assert.Equal(t, 200, int(responseBody["code"].(float64)))
// 		assert.Equal(t, "Succeed", responseBody["message"])
// 		// assert.Equal(t, 1, responseBody["data"].(map[string]interface{})["id"])
// 		assert.Equal(t, "Fani", responseBody["data"].(map[string]interface{})["name"])
// 		assert.Equal(t, "fani@aol.com", responseBody["data"].(map[string]interface{})["email"])

// 	}
// }
