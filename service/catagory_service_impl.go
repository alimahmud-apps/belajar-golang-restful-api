package service

import (
	"belajar-golang-restful-api/exception"
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/api"
	"belajar-golang-restful-api/model/domain"
	"belajar-golang-restful-api/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type CatagoryServiceImpl struct {
	CatagoryRepository repository.CatagoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCatagoryService(catagoryRepository repository.CatagoryRepository, DB *sql.DB, validate *validator.Validate) CatagoryService {
	return &CatagoryServiceImpl{
		CatagoryRepository: catagoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CatagoryServiceImpl) Create(ctx context.Context, request api.CatagoryCreateRequest) api.CatagoryResponse {
	errVal := service.Validate.Struct(request)
	helper.PanicIfError(errVal)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	catagory := domain.Catagory{
		Name: request.Name,
	}

	catagory = service.CatagoryRepository.Save(ctx, tx, catagory)

	return helper.ToCatagoryResponse(catagory)
}

func (service *CatagoryServiceImpl) Update(ctx context.Context, request api.CatagoryUpdateRequest) api.CatagoryResponse {
	errVal := service.Validate.Struct(request)
	helper.PanicIfError(errVal)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	catagory, err := service.CatagoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	catagory.Name = request.Name

	catagory = service.CatagoryRepository.Update(ctx, tx, catagory)

	return helper.ToCatagoryResponse(catagory)
}

func (service *CatagoryServiceImpl) Delete(ctx context.Context, catagoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	catagory, err := service.CatagoryRepository.FindById(ctx, tx, catagoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CatagoryRepository.Delete(ctx, tx, catagory)
}

func (service *CatagoryServiceImpl) FindById(ctx context.Context, catagoryId int) api.CatagoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	catagory, err := service.CatagoryRepository.FindById(ctx, tx, catagoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCatagoryResponse(catagory)
}

func (service *CatagoryServiceImpl) FindAll(ctx context.Context) []api.CatagoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	catagories := service.CatagoryRepository.FindAll(ctx, tx)

	return helper.ToCatagoryResponses(catagories)
}
