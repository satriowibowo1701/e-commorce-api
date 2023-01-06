package service

import (
	"context"
	"errors"

	"github.com/satriowibowo1701/e-commorce-api/model"
)

func (service *InitService) Create(ctx context.Context, request model.ProdukRequest) error {

	err := service.Validate.Struct(request)
	if err != nil {
		return err
	}
	tx2, _ := service.DB.Begin()
	_, err2 := service.ProdukRepostory.FindByName(ctx, tx2, request.Name)
	if err2 == nil {
		return errors.New("name already exists")
	}
	tx, _ := service.DB.Begin()
	_ = service.ProdukRepostory.Create(ctx, tx, request)

	return nil

}

func (service *InitService) Update(ctx context.Context, request model.ProdukUpdate) error {

	err := service.Validate.Struct(request)
	if err != nil {
		return err
	}

	tx, _ := service.DB.Begin()

	_, err1 := service.ProdukRepostory.FindById(ctx, tx, int(request.Product_id))
	if err1 != nil {
		return err1
	}

	err2 := service.ProdukRepostory.Update(ctx, tx, request)
	if err2 != nil {
		return err2
	}

	return nil
}

func (service *InitService) Delete(ctx context.Context, produkid int) error {

	tx, _ := service.DB.Begin()
	_, err2 := service.ProdukRepostory.FindById(ctx, tx, int(produkid))
	if err2 != nil {
		return errors.New("produk not found")
	}
	err1 := service.ProdukRepostory.DeleteById(ctx, tx, int64(produkid))
	if err1 != nil {
		return err1
	}
	return nil
}

func (service *InitService) FindById(ctx context.Context, productid int) (*model.Produk, error) {
	tx, _ := service.DB.Begin()
	produk, err := service.ProdukRepostory.FindById(ctx, tx, productid)
	if err != nil {
		return nil, errors.New("produk not found")
	}

	return produk, nil
}

func (service *InitService) FindAll(ctx context.Context) ([]*model.Produk, error) {
	tx, _ := service.DB.Begin()

	products, err := service.ProdukRepostory.FindAll(ctx, tx)
	if err != nil {
		return nil, err
	}
	return products, nil

}
