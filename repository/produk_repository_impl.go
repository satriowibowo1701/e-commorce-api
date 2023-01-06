package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/satriowibowo1701/e-commorce-api/helper"
	"github.com/satriowibowo1701/e-commorce-api/model"
)

type ProductImpl struct {
}

func NewProductRepo() ProductRepo {
	return &ProductImpl{}
}

func (repository *ProductImpl) Create(ctx context.Context, tx *sql.Tx, produk model.ProdukRequest) error {
	SQL := "insert into product(name,quantity,price) values (?,?,?)"
	_, err := tx.ExecContext(ctx, SQL, produk.Name, produk.Qty, produk.Price)
	return helper.TxRollback(err, tx, "error Creating Produk")
}

func (repository *ProductImpl) Update(ctx context.Context, tx *sql.Tx, produk model.ProdukUpdate) error {
	SQL := "update product set name = ?,quantity=?,price=? WHERE product_id = ?"
	_, err := tx.ExecContext(ctx, SQL, produk.Name, produk.Qty, produk.Price, produk.Product_id)

	return helper.TxRollback(err, tx, "Error updating Produk")
}

func (repository *ProductImpl) FindById(ctx context.Context, tx *sql.Tx, produkId int) (*model.Produk, error) {
	SQL := "select product_id,name,quantity from product where product_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, produkId)
	if err != nil {
		return nil, errors.New("Error Sql")
	}
	defer rows.Close()
	defer tx.Commit()
	produk := model.Produk{}
	if rows.Next() {
		err := rows.Scan(&produk.Product_id, &produk.Name, &produk.Qty)
		if err != nil {
			return nil, errors.New("error Scan")
		}
		return &produk, nil
	} else {
		return nil, errors.New("Produk not found")
	}
}

func (repository *ProductImpl) FindByName(ctx context.Context, tx *sql.Tx, name string) (*model.Produk, error) {
	SQL := "select product_id,name,quantity from product where name = ?"
	rows, err := tx.QueryContext(ctx, SQL, name)
	if err != nil {
		return nil, errors.New("Error Sql")
	}
	defer rows.Close()
	defer tx.Commit()
	produk := model.Produk{}
	if rows.Next() {
		err := rows.Scan(&produk.Product_id, &produk.Name, &produk.Qty)
		if err != nil {
			return nil, errors.New("error Scan")
		}
		return &produk, nil
	} else {
		return nil, errors.New("Produk not found")
	}
}

func (repository *ProductImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]*model.Produk, error) {
	SQL := "select product_id,name,quantity from product"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	defer tx.Commit()
	var products []*model.Produk
	for rows.Next() {
		product := model.Produk{}
		err := rows.Scan()
		if err != nil {
			return nil, errors.New("Cannot Scaning")
		}
		products = append(products, &product)
	}
	return products, nil
}

func (repository *ProductImpl) DeleteById(ctx context.Context, tx *sql.Tx, productid int64) error {
	SQL := "DELETE FROM product WHERE product_id =?"
	_, err := tx.ExecContext(ctx, SQL, productid)
	return helper.TxRollback(err, tx, "Error Deleting Produk")
}
