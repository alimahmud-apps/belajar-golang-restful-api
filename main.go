package main

import (
	"belajar-golang-restful-api/app"
	"belajar-golang-restful-api/controller"
	"belajar-golang-restful-api/exception"
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/middleware"
	"belajar-golang-restful-api/repository"
	"belajar-golang-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	repository := repository.NewCatagoryRepository()
	service := service.NewCatagoryService(repository, db, validate)
	controller := controller.NewCatagoryController(service)

	router := httprouter.New()

	router.GET("/api/v1/catagories", controller.FindAll)
	router.GET("/api/v1/catagories/:catagoryId", controller.FindById)
	router.POST("/api/v1/catagories", controller.Create)
	router.PUT("/api/v1/catagories/:catagoryId", controller.Update)
	router.DELETE("/api/v1/catagories/:catagoryId", controller.Delete)

	router.PanicHandler = exception.ErrorHandler
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
