package repository

import (
	"context"
	"database/sql"

	"github.com/satriowibowo1701/e-commorce-api/model"
)

type ProductRepo interface {
	Create(ctx context.Context, tx *sql.Tx, produk model.ProdukRequest) error
	Update(ctx context.Context, tx *sql.Tx, produk model.Produk) error
	FindById(ctx context.Context, tx *sql.Tx, produkId int) (*model.Produk, error)
	FindByName(ctx context.Context, tx *sql.Tx, name string) (*model.Produk, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]*model.Produk, error)
	DeleteById(ctx context.Context, tx *sql.Tx, productid int64) error
}
