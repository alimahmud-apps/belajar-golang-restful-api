package exception

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/api"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writer, request, err) {
		return
	}
	if validationError(writer, request, err) {
		return
	}
	internalServerError(writer, request, err)
}
func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		apiResponse := api.ApiRespone{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}
		helper.WriterToResponseBody(writer, apiResponse)
		return true
	} else {
		return false
	}

}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		apiResponse := api.ApiRespone{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}
		helper.WriterToResponseBody(writer, apiResponse)
		return true
	} else {
		return false
	}

}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)
	apiResponse := api.ApiRespone{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}
	helper.WriterToResponseBody(writer, apiResponse)
}
