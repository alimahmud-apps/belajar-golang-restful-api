package controller

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/api"
	"belajar-golang-restful-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CatagoryControllerImpl struct {
	CatagoryService service.CatagoryService
}

func NewCatagoryController(catagoryService service.CatagoryService) CatagoryController {
	return &CatagoryControllerImpl{
		CatagoryService: catagoryService,
	}
}

func (controller *CatagoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	catagoryCreateRequest := api.CatagoryCreateRequest{}
	helper.ReadFromRequestBody(request, &catagoryCreateRequest)

	catagoryResponse := controller.CatagoryService.Create(request.Context(), catagoryCreateRequest)
	apiResponse := api.ApiRespone{
		Code:   200,
		Status: "OK",
		Data:   catagoryResponse,
	}

	helper.WriterToResponseBody(writer, apiResponse)
}

func (controller *CatagoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	catagoryUpdateRequest := api.CatagoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &catagoryUpdateRequest)

	catagoryId := params.ByName("catagoryId")
	id, err := strconv.Atoi(catagoryId)
	helper.PanicIfError(err)

	catagoryUpdateRequest.Id = id

	catagoryResponse := controller.CatagoryService.Update(request.Context(), catagoryUpdateRequest)
	apiResponse := api.ApiRespone{
		Code:   200,
		Status: "OK",
		Data:   catagoryResponse,
	}

	helper.WriterToResponseBody(writer, apiResponse)
}

func (controller *CatagoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	catagoryId := params.ByName("catagoryId")
	id, err := strconv.Atoi(catagoryId)
	helper.PanicIfError(err)

	controller.CatagoryService.Delete(request.Context(), id)
	apiResponse := api.ApiRespone{
		Code:   200,
		Status: "OK",
	}

	helper.WriterToResponseBody(writer, apiResponse)
}

func (controller *CatagoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	catagoryId := params.ByName("catagoryId")
	id, err := strconv.Atoi(catagoryId)
	helper.PanicIfError(err)

	catagoryResponse := controller.CatagoryService.FindById(request.Context(), id)
	apiResponse := api.ApiRespone{
		Code:   200,
		Status: "OK",
		Data:   catagoryResponse,
	}

	helper.WriterToResponseBody(writer, apiResponse)
}

func (controller *CatagoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	catagoryResponses := controller.CatagoryService.FindAll(request.Context())
	apiResponse := api.ApiRespone{
		Code:   200,
		Status: "OK",
		Data:   catagoryResponses,
	}

	helper.WriterToResponseBody(writer, apiResponse)
}
