package test

import (
	"belajar-golang-restful-api/app"
	"belajar-golang-restful-api/controller"
	"belajar-golang-restful-api/exception"
	"belajar-golang-restful-api/middleware"
	"belajar-golang-restful-api/model/domain"
	"belajar-golang-restful-api/repository"
	"belajar-golang-restful-api/service"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = ""
	dbname   = "belajar_golang_restful_api_test"
)

func setupTestDB() *sql.DB {
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		user, password, host, port, dbname)

	db, err := sql.Open("mysql", mysqlInfo)

	if err != nil {
		panic(err.Error())
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func setupRourer(db *sql.DB) http.Handler {

	validate := validator.New()
	repository := repository.NewCatagoryRepository()
	service := service.NewCatagoryService(repository, db, validate)
	controller := controller.NewCatagoryController(service)
	router := app.NewRouter(controller)
	router.PanicHandler = exception.ErrorHandler
	return middleware.NewAuthMiddleware(router)
}

func truncateCatagory(db *sql.DB) {
	db.Exec("TRUNCATE catagory")
}
func TestCreateCatagorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCatagory(db)
	router := setupRourer(db)

	requestBody := strings.NewReader(`{"name":"Laptop"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/catagories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("x-api-key", "rahasia")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// fmt.Println(responseBody)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Laptop", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateCatagoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCatagory(db)
	router := setupRourer(db)

	requestBody := strings.NewReader(`{"name":""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/catagories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("x-api-key", "rahasia")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestUpdateCatagorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCatagory(db)
	tx, _ := db.Begin()
	catagoryRepository := repository.NewCatagoryRepository()
	catagory := catagoryRepository.Save(context.Background(), tx, domain.Catagory{
		Name: "MAINAN",
	})
	tx.Commit()

	router := setupRourer(db)

	id := strconv.Itoa(catagory.Id)
	requestBody := strings.NewReader(`{"name":"Laptop"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/v1/catagories/"+id, requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("x-api-key", "rahasia")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// fmt.Println(responseBody)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, catagory.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Laptop", responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateCatagoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCatagory(db)
	tx, _ := db.Begin()
	catagoryRepository := repository.NewCatagoryRepository()
	catagory := catagoryRepository.Save(context.Background(), tx, domain.Catagory{
		Name: "MAINAN",
	})
	tx.Commit()

	router := setupRourer(db)

	id := strconv.Itoa(catagory.Id)
	requestBody := strings.NewReader(`{"name":""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/v1/catagories/"+id, requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("x-api-key", "rahasia")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestGetCatagorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCatagory(db)
	tx, _ := db.Begin()
	catagoryRepository := repository.NewCatagoryRepository()
	catagory := catagoryRepository.Save(context.Background(), tx, domain.Catagory{
		Name: "MAINAN",
	})
	tx.Commit()

	router := setupRourer(db)

	id := strconv.Itoa(catagory.Id)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/catagories/"+id, nil)
	request.Header.Add("x-api-key", "rahasia")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// fmt.Println(responseBody)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, catagory.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, catagory.Name, responseBody["data"].(map[string]interface{})["name"])
}

func TestGetCatagoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCatagory(db)
	router := setupRourer(db)

	id := "404"

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/catagories/"+id, nil)
	request.Header.Add("x-api-key", "rahasia")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// fmt.Println(responseBody)
	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestDeleteCatagorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCatagory(db)
	tx, _ := db.Begin()
	catagoryRepository := repository.NewCatagoryRepository()
	catagory := catagoryRepository.Save(context.Background(), tx, domain.Catagory{
		Name: "MAINAN",
	})
	tx.Commit()

	router := setupRourer(db)

	id := strconv.Itoa(catagory.Id)
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/v1/catagories/"+id, nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("x-api-key", "rahasia")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// fmt.Println(responseBody)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestDeleteCatagoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCatagory(db)
	router := setupRourer(db)

	id := "404"
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/v1/catagories/"+id, nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("x-api-key", "rahasia")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// fmt.Println(responseBody)
	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestListGetCatagory(t *testing.T) {
	db := setupTestDB()
	truncateCatagory(db)
	tx, _ := db.Begin()
	catagoryRepository := repository.NewCatagoryRepository()
	catagory1 := catagoryRepository.Save(context.Background(), tx, domain.Catagory{
		Name: "MAINAN",
	})
	catagory2 := catagoryRepository.Save(context.Background(), tx, domain.Catagory{
		Name: "LAPTOP",
	})
	tx.Commit()

	router := setupRourer(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/catagories", nil)
	request.Header.Add("x-api-key", "rahasia")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// fmt.Println(responseBody["data"].([]interface{})[0].(map[string]interface{})["id"], catagory1.Id, catagory2.Id)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	responseCatagory1 := responseBody["data"].([]interface{})[0].(map[string]interface{})
	responseCatagory2 := responseBody["data"].([]interface{})[1].(map[string]interface{})

	assert.Equal(t, catagory1.Id, int(responseCatagory1["id"].(float64)))
	assert.Equal(t, catagory1.Name, responseCatagory1["name"])

	assert.Equal(t, catagory2.Id, int(responseCatagory2["id"].(float64)))
	assert.Equal(t, catagory2.Name, responseCatagory2["name"])
}

func TestUnauthorize(t *testing.T) {
	db := setupTestDB()
	truncateCatagory(db)

	router := setupRourer(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/catagories", nil)
	request.Header.Add("x-api-key", "salah")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])

}
