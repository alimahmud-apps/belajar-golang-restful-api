package main

import (
	"belajar-golang-restful-api/app"
	"belajar-golang-restful-api/controller"
	"belajar-golang-restful-api/exception"
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/middleware"
	"belajar-golang-restful-api/repository"
	"belajar-golang-restful-api/service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port env is required")
	}
	db := app.NewDB()
	validate := validator.New()
	repository := repository.NewCatagoryRepository()
	service := service.NewCatagoryService(repository, db, validate)
	controller := controller.NewCatagoryController(service)

	router := app.NewRouter(controller)

	router.PanicHandler = exception.ErrorHandler
	server := http.Server{
		Addr:    ":" + port,
		Handler: middleware.NewAuthMiddleware(router),
	}
	fmt.Println("Server running in localhost:" + port)
	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
