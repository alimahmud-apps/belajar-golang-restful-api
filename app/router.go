package app

import (
	"belajar-golang-restful-api/controller"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(controller controller.CatagoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/v1/catagories", controller.FindAll)
	router.GET("/api/v1/catagories/:catagoryId", controller.FindById)
	router.POST("/api/v1/catagories", controller.Create)
	router.PUT("/api/v1/catagories/:catagoryId", controller.Update)
	router.DELETE("/api/v1/catagories/:catagoryId", controller.Delete)

	return router
}
