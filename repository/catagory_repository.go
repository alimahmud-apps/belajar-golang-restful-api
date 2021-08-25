package repository

import (
	"belajar-golang-restful-api/model/domain"
	"context"
	"database/sql"
)

type CatagoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, catagory domain.Catagory) domain.Catagory
	Update(ctx context.Context, tx *sql.Tx, catagory domain.Catagory) domain.Catagory
	Delete(ctx context.Context, tx *sql.Tx, catagory domain.Catagory)
	FindById(ctx context.Context, tx *sql.Tx, catagoryId int) (domain.Catagory, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Catagory
}
