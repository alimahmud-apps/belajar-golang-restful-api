package service

import (
	"belajar-golang-restful-api/model/api"
	"context"
)

type CatagoryService interface {
	Create(ctx context.Context, request api.CatagoryCreateRequest) api.CatagoryResponse
	Update(ctx context.Context, request api.CatagoryUpdateRequest) api.CatagoryResponse
	Delete(ctx context.Context, catagoryId int)
	FindById(ctx context.Context, catagoryId int) api.CatagoryResponse
	FindAll(ctx context.Context) []api.CatagoryResponse
}
