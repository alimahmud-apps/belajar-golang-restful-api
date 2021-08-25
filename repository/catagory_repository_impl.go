package repository

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/domain"
	"context"
	"database/sql"
	"errors"
)

type CatagoryRepositoryImpl struct {
}

func NewCatagoryRepository() CatagoryRepository {
	return &CatagoryRepositoryImpl{}
}

func (repository *CatagoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, catagory domain.Catagory) domain.Catagory {
	stringSQL := "insert into catagory(name) values(?)"
	result, err := tx.ExecContext(ctx, stringSQL, catagory.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	catagory.Id = int(id)
	return catagory

}

func (repository *CatagoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, catagory domain.Catagory) domain.Catagory {
	stringSQL := "update catagory set name=? where id=?"
	_, err := tx.ExecContext(ctx, stringSQL, catagory.Name, catagory.Id)
	helper.PanicIfError(err)

	return catagory
}

func (repository *CatagoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, catagory domain.Catagory) {
	stringSQL := "delete from catagory where id = ?"
	_, err := tx.ExecContext(ctx, stringSQL, catagory.Id)
	helper.PanicIfError(err)
}

func (repository *CatagoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, catagoryId int) (domain.Catagory, error) {
	stringSQL := "select id,name from catagory where id=?"
	rows, err := tx.QueryContext(ctx, stringSQL, catagoryId)
	helper.PanicIfError(err)
	defer rows.Close()
	catagory := domain.Catagory{}
	if rows.Next() {
		err = rows.Scan(&catagory.Id, &catagory.Name)
		helper.PanicIfError(err)
		return catagory, nil
	} else {
		return catagory, errors.New("catagory not found")
	}
}

func (repository *CatagoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Catagory {
	stringSQL := "select id,name from catagory "
	rows, err := tx.QueryContext(ctx, stringSQL)
	helper.PanicIfError(err)
	defer rows.Close()
	var catagories []domain.Catagory
	for rows.Next() {
		catagory := domain.Catagory{}
		rows.Scan(&catagory.Id, &catagory.Name)
		catagories = append(catagories, catagory)
	}

	return catagories
}
