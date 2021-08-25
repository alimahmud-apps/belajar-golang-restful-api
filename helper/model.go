package helper

import (
	"belajar-golang-restful-api/model/api"
	"belajar-golang-restful-api/model/domain"
)

func ToCatagoryResponse(catagory domain.Catagory) api.CatagoryResponse {
	return api.CatagoryResponse{
		Id:   catagory.Id,
		Name: catagory.Name,
	}
}

func ToCatagoryResponses(catagories []domain.Catagory) []api.CatagoryResponse {
	var catagoryResponses []api.CatagoryResponse
	for _, catagory := range catagories {
		catagoryResponses = append(catagoryResponses, ToCatagoryResponse(catagory))
	}

	return catagoryResponses
}
