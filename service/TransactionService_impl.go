package service

import (
	"context"
	"errors"

	"github.com/satriowibowo1701/e-commorce-api/model"
)

func (service *InitService) CreateTransaction(ctx context.Context, request model.TransactionRequest) error {

	err := service.Validate.Struct(request)
	if err != nil {
		return err
	}
	if len(request.OrderItems) == 0 {
		return errors.New("no order items")
	}
	prices := make(chan int)
	newprice := 0
	for _, items := range request.OrderItems {
		go func(price int) {
			prices <- price
		}(int(items.OrderPrice))
		newprice += <-prices
	}
	request.Total = int64(newprice)
	tx, _ := service.DB.Begin()
	err, id := service.TransactionRepository.CreateTransaction(ctx, tx, request)
	if err != nil {
		return err
	}
	txdelete, _ := service.DB.Begin()
	go service.TransactionRepository.DeleteTempTransaction(context.Background(), txdelete, request.CustomerId)
	for _, items := range request.OrderItems {
		tx2, _ := service.DB.Begin()
		go func(items *model.OrderItem) {
			service.TransactionRepository.InsertOrderItems(context.Background(), tx2, items, int64(id))
		}(items)
	}

	return nil
}

func (service *InitService) UpdateTmpTransaction(ctx context.Context, request model.TempUpdateTransactionRequest) error {

	err := service.Validate.Struct(request)
	if err != nil {
		return err
	}
	tx, _ := service.DB.Begin()

	go service.TransactionRepository.UpdateTempTransaction(ctx, tx, request)
	err2 := service.TransactionRepository.GetTempTransactionsByid(ctx, tx, request.Id)
	if err2 != nil {
		return err2

	}
	return nil

}

func (service *InitService) DeleteTmpTransaction(ctx context.Context, idtemptrx int64, cusid int64) error {

	tx, _ := service.DB.Begin()

	if idtemptrx == -1 || cusid == -1 {
		return errors.New("Id/cusid not attached")
	}

	err1 := service.TransactionRepository.DeleteTempTransactionByid(ctx, tx, idtemptrx, cusid)
	if err1 != nil {
		return err1
	}
	return nil
}

func (service *InitService) FindAllTmpTransaction(ctx context.Context, cusid int64) ([]*model.TempTransaction, error) {
	if cusid == -1 {
		return nil, errors.New("No Cookie Id Found")
	}
	tx, _ := service.DB.Begin()

	items, err := service.TransactionRepository.GetAllTempTransactionsCus(ctx, tx, cusid)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (service *InitService) FindAllTmpTransactionCustomer(ctx context.Context, cusid int64) ([]*model.TempTransaction, error) {
	if cusid == -1 {
		return nil, errors.New("No Cookie Id Found")
	}
	tx, _ := service.DB.Begin()

	items, err := service.TransactionRepository.GetAllTempTransactionsCus(ctx, tx, cusid)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (service *InitService) FindAllTransactionCustomer(ctx context.Context) ([]*model.TransactionAdmin, error) {
	tx, _ := service.DB.Begin()
	items, err := service.TransactionRepository.GetAllTransaction(ctx, tx)
	if err != nil {
		return nil, err
	}

	res := make(chan *model.TransactionAdmin)
	defer close(res)
	newres := []*model.TransactionAdmin{}
	for _, itemsss := range items {
		tx2, _ := service.DB.Begin()
		ctx2 := context.Background()
		go func(itemss *model.TransactionAdmin, store []*model.TransactionAdmin) {
			orderitems := service.TransactionRepository.GetAllOrderItems(ctx2, tx2, itemss.TransactionId)
			newtrx := &model.TransactionAdmin{
				TransactionId:    itemss.TransactionId,
				CustomerName:     itemss.CustomerName,
				CustomerId:       itemss.CustomerId,
				Date:             itemss.Date,
				Status:           itemss.Status,
				TransactionTotal: itemss.TransactionTotal,
				OrderItem:        orderitems,
			}
			res <- newtrx
		}(itemsss, newres)
		newres = append(newres, <-res)
	}

	return newres, nil
}

func (service *InitService) FindAllTransactionByStatus(ctx context.Context, status int, cusid int) ([]*model.TransactionCus, error) {
	if cusid == -1 {
		return nil, errors.New("No Cookie Id Found")
	}
	tx, _ := service.DB.Begin()
	items, err := service.TransactionRepository.GetAllTransactionsByStatusCus(ctx, tx, int64(status), int64(cusid))
	if err != nil {
		return nil, err
	}
	res := make(chan *model.TransactionCus)
	defer close(res)
	newres := []*model.TransactionCus{}
	for _, itemsss := range items {
		tx2, _ := service.DB.Begin()
		ctx2 := context.Background()
		go func(itemss *model.TransactionCus, store []*model.TransactionCus) {
			orderitems := service.TransactionRepository.GetAllOrderItems(ctx2, tx2, itemss.TransactionId)
			newtrx := &model.TransactionCus{
				TransactionId: itemss.TransactionId,
				CustomerId:    itemss.CustomerId,
				Date:          itemss.Date,
				Status:        itemss.Status,
				OrderItem:     orderitems,
			}
			res <- newtrx
		}(itemsss, newres)

		newres = append(newres, <-res)
	}

	return newres, nil
}

func (service *InitService) FindAllTransactionById(ctx context.Context, cusid int) ([]*model.TransactionCus, error) {
	if cusid == -1 {
		return nil, errors.New("No Cookie Id Found")
	}
	tx, _ := service.DB.Begin()
	items, err := service.TransactionRepository.GetAllTransactionById(ctx, tx, int64(cusid))
	if err != nil {
		return nil, err
	}
	res := make(chan *model.TransactionCus)
	defer close(res)
	newres := []*model.TransactionCus{}
	for _, itemsss := range items {
		tx2, _ := service.DB.Begin()
		ctx2 := context.Background()
		go func(itemss *model.TransactionCus, store []*model.TransactionCus) {
			orderitems := service.TransactionRepository.GetAllOrderItems(ctx2, tx2, itemss.TransactionId)
			newtrx := &model.TransactionCus{
				TransactionId: itemss.TransactionId,
				CustomerId:    itemss.CustomerId,
				Date:          itemss.Date,
				Status:        itemss.Status,
				OrderItem:     orderitems,
			}
			res <- newtrx
		}(itemsss, newres)

		newres = append(newres, <-res)
	}

	return newres, nil
}
func (service *InitService) InsertTmpTransaction(ctx context.Context, req model.TempTransactionRequest) error {
	err := service.Validate.Struct(req)
	if err != nil {
		return err
	}
	tx, _ := service.DB.Begin()
	service.TransactionRepository.InsertTempTransaction(ctx, tx, req)

	return nil
}

func (service *InitService) FindAllTrxByTransactionid(ctx context.Context, trxid int64, cusid int64) (*model.TransactionCus, error) {

	if cusid == -1 {
		return nil, errors.New("No Cookie Id Found")
	}
	tx, _ := service.DB.Begin()
	items, err := service.TransactionRepository.GetAllTransactionsByTransactionid(ctx, tx, trxid, cusid)
	if err != nil {
		return nil, err
	}
	return items, nil

}
