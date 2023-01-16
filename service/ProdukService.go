package service

import (
	"context"
	"errors"

	"github.com/satriowibowo1701/e-commorce-api/helper"
	"github.com/satriowibowo1701/e-commorce-api/model"
)

func (service *InitService) Create(ctx context.Context, request model.ProdukRequest) error {

	err := service.Validate.Struct(request)
	if err != nil {
		return err
	}
	tx, _ := service.DB.Begin()
	_, err2 := service.ProdukRepostory.FindByName(ctx, tx, request.Name)
	if err2 == nil {
		return errors.New("name already exists")
	}
	_ = service.ProdukRepostory.Create(ctx, tx, request)
	tx.Commit()
	return nil

}

func (service *InitService) Update(ctx context.Context, request model.ProdukUpdate) error {

	errchname := make(chan error)
	errchid := make(chan error)
	defer close(errchid)
	defer close(errchname)
	go func() {
		tx2, _ := service.DB.Begin()
		defer tx2.Commit()
		_, err1 := service.ProdukRepostory.FindById(context.Background(), tx2, int(request.Product_id))
		if err1 != nil {
			errchid <- err1
			return
		}
		errchid <- nil

	}()
	go func() {
		tx1, _ := service.DB.Begin()
		defer tx1.Commit()
		_, err2 := service.ProdukRepostory.FindByName(context.Background(), tx1, request.Name)
		errchname <- err2

	}()
	err := service.Validate.Struct(request)
	err1 := <-errchid
	err2 := <-errchname
	if err != nil {
		return err
	}
	if err1 != nil {
		return err1
	}
	if err2 == nil {
		return errors.New("Product Name Already Exist")
	}
	tx, _ := service.DB.Begin()
	service.ProdukRepostory.Update(ctx, tx, request)
	tx.Commit()

	return nil
}

func (service *InitService) Delete(ctx context.Context, produkid int) error {

	tx, _ := service.DB.Begin()
	_, err2 := service.ProdukRepostory.FindById(ctx, tx, int(produkid))
	defer helper.TxRollback(err2, tx, "Error")
	if err2 != nil {
		return errors.New("produk not found")
	}
	service.ProdukRepostory.DeleteById(ctx, tx, int64(produkid))
	return nil
}

func (service *InitService) FindById(ctx context.Context, productid int) (*model.Produk, error) {
	tx, _ := service.DB.Begin()
	produk, err := service.ProdukRepostory.FindById(ctx, tx, productid)
	defer helper.TxRollback(err, tx, "Error")
	if err != nil {
		return nil, errors.New("produk not found")
	}

	return produk, nil
}

func (service *InitService) FindAll(ctx context.Context) ([]*model.Produk, error) {
	tx, _ := service.DB.Begin()
	products, err := service.ProdukRepostory.FindAll(ctx, tx)
	defer helper.TxRollback(err, tx, "Error")
	if err != nil {
		return nil, err
	}
	return products, nil

}

func (service *InitService) FindAllPrdkAdmin(ctx context.Context) ([]*model.Produk, error) {
	tx, _ := service.DB.Begin()
	products, err := service.ProdukRepostory.FindAllAdmin(ctx, tx)
	defer helper.TxRollback(err, tx, "Error")
	if err != nil {
		return nil, err
	}
	return products, nil

}
