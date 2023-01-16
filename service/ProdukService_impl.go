package service

import (
	"context"

	"github.com/satriowibowo1701/e-commorce-api/model"
)

type ProductService interface {
	Create(ctx context.Context, request model.ProdukRequest) error
	Update(ctx context.Context, request model.ProdukUpdate) error
	Delete(ctx context.Context, produkid int) error
	FindById(ctx context.Context, productid int) (*model.Produk, error)
	FindAll(ctx context.Context) ([]*model.Produk, error)
	FindAllPrdkAdmin(ctx context.Context) ([]*model.Produk, error)
}
