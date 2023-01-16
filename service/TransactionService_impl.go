package service

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"sync"

	"github.com/satriowibowo1701/e-commorce-api/helper"
	"github.com/satriowibowo1701/e-commorce-api/model"
)

func (service *InitService) CreateTransaction(ctx context.Context, request model.TransactionRequest, csid int64) error {

	err := service.Validate.Struct(request)
	if err != nil {
		return err
	}
	if len(request.OrderItems) == 0 {
		return errors.New("no order items")
	}
	txpool := sync.Pool{
		New: func() interface{} {
			db, _ := service.DB.Begin()
			return db
		},
	}
	var newprice int64
	for _, items := range request.OrderItems {
		newprice += items.OrderPrice * items.OrderQty
	}
	request.Total = newprice
	request.CustomerId = csid
	tx := txpool.Get().(*sql.Tx)
	id, err2 := service.TransactionRepository.CreateTransaction(ctx, tx, request)
	defer helper.TxRollback(err2, tx, "Error Create Trx")
	if err2 != nil {
		return err2
	}
	wg := sync.WaitGroup{}
	limit := make(chan struct{}, 20)
	defer close(limit)
	go service.TransactionRepository.DeleteTempTransaction(ctx, tx, request.CustomerId)
	for _, items := range request.OrderItems {
		wg.Add(1)
		limit <- struct{}{}
		go func(itemss *model.OrderItem) {
			defer wg.Done()
			defer func() { <-limit }()
			tx3 := txpool.Get().(*sql.Tx)
			service.TransactionRepository.InsertOrderItems(ctx, tx, itemss, int64(id))
			data, _ := service.ProdukRepostory.FindById(context.Background(), tx3, int(itemss.ProductId))
			tx3.Commit()
			service.ProdukRepostory.UpdateQty(ctx, tx, data.Qty-itemss.OrderQty, itemss.ProductId)
		}(items)
	}

	wg.Wait()

	return nil
}

func (service *InitService) UpdateTmpTransaction(ctx context.Context, request model.TempUpdateTransactionRequest) error {
	err := service.Validate.Struct(request)
	if err != nil {
		return err
	}
	tx, _ := service.DB.Begin()

	err1 := service.TransactionRepository.UpdateTempTransaction(ctx, tx, request)

	return helper.TxRollback(err1, tx, "Error Updating TmpTransaction")

}

func (service *InitService) DeleteTmpTransaction(ctx context.Context, idtemptrx int64, cusid int64) error {

	tx, _ := service.DB.Begin()
	if idtemptrx == -1 || cusid == -1 {
		return errors.New("Id/cusid not attached")
	}
	err1 := service.TransactionRepository.DeleteTempTransactionByid(ctx, tx, idtemptrx, cusid)
	defer helper.TxRollback(err1, tx, "error")
	if err1 != nil {
		return err1
	}
	return nil
}

func (service *InitService) FindAllTmpTransactionCustomer(ctx context.Context, cusid int64) ([]*model.TempTransaction, error) {
	if cusid == -1 {
		return nil, errors.New("No Cookie Id Found")
	}
	tx, _ := service.DB.Begin()

	items, err := service.TransactionRepository.GetAllTempTransactionsCus(ctx, tx, cusid)
	defer helper.TxRollback(err, tx, "error")
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (service *InitService) FindAllTransactionCustomer(ctx context.Context) ([]model.TransactionAdmin, error) {
	txpool := sync.Pool{
		New: func() interface{} {
			db, _ := service.DB.Begin()
			return db
		},
	}
	tx := txpool.Get().(*sql.Tx)
	items, err := service.TransactionRepository.GetAllTransaction(ctx, tx)
	tx.Commit()
	if err != nil {
		return nil, err
	}
	if cache := service.CacheData[-1]; cache != nil {
		data, _ := cache.([]model.TransactionAdmin)
		if len(data) == len(items) {
			return data, nil
		}
	}
	wg := sync.WaitGroup{}
	limit := make(chan struct{}, 30)
	for index, itemsss := range items {
		wg.Add(1)
		limit <- struct{}{}
		go func(itemss model.TransactionAdmin, index int) {
			defer wg.Done()
			defer func() { <-limit }()
			tx2 := txpool.Get().(*sql.Tx)
			defer tx2.Commit()
			items[index].OrderItem = service.TransactionRepository.GetAllOrderItems(context.Background(), tx2, itemss.TransactionId)
		}(itemsss, index)
	}
	wg.Wait()
	if cache := service.CacheData[-1]; cache == nil || cache != nil {
		if cache != nil {
			data, _ := cache.([]model.TransactionAdmin)
			if len(data) != len(items) {
				service.CacheData[-1] = items
				return items, nil
			}
		}
		service.CacheData[-1] = items
	}
	return items, nil
}

func (service *InitService) FindAllTransactionByStatus(ctx context.Context, status int, cusid int) ([]model.TransactionCus, error) {
	if cusid == -1 {
		return nil, errors.New("No Cookie Id Found")
	}
	txpool := sync.Pool{
		New: func() interface{} {
			db, _ := service.DB.Begin()
			return db
		},
	}
	tx, _ := txpool.Get().(*sql.Tx)
	items, err := service.TransactionRepository.GetAllTransactionsByStatusCus(ctx, tx, int64(status), int64(cusid))
	defer helper.TxRollback(err, tx, "Error Get Transaction")
	if err != nil {
		return nil, err
	}
	if cache := service.CacheData[int(cusid)]; cache != nil {
		data, _ := cache.([]model.TransactionCus)
		if len(data) == len(items) {
			return data, nil
		}
	}
	wg := sync.WaitGroup{}
	limit := make(chan struct{}, 30)
	defer close(limit)
	for index, itemsss := range items {
		wg.Add(1)
		limit <- struct{}{}
		go func(itemss model.TransactionCus, index int64) {
			defer func() { <-limit }()
			defer wg.Done()
			tx2 := txpool.Get().(*sql.Tx)
			defer tx2.Commit()
			ctx2 := context.Background()
			paymentInfo, _ := service.PaymentsRepository.GetAllPaymentByid(ctx2, tx2, itemss.PaymentId)
			items[index].PaymentInfo = paymentInfo
			orderitemss := service.TransactionRepository.GetAllOrderItems(ctx2, tx2, itemss.TransactionId)
			items[index].OrderItem = orderitemss
		}(itemsss, int64(index))
	}
	wg.Wait()
	if service.CacheData[int(cusid)] == nil || service.CacheData[int(cusid)] != nil {
		if service.CacheData[int(cusid)] != nil {
			cache := service.CacheData[int(cusid)].([]model.TransactionCus)
			if len(cache) != len(items) {
				service.CacheData[int(cusid)] = items
				return items, nil
			}
		}
		service.CacheData[int(cusid)] = items
	}
	return items, nil
}

func (service *InitService) FindAllTransactionById(ctx context.Context, cusid int) ([]model.TransactionCus, error) {
	if cusid == -1 {
		return nil, errors.New("No Cookie Id Found")
	}
	txpool := sync.Pool{
		New: func() interface{} {
			db, _ := service.DB.Begin()
			return db
		},
	}
	tx := txpool.Get().(*sql.Tx)
	items, err := service.TransactionRepository.GetAllTransactionById(ctx, tx, int64(cusid))
	defer helper.TxRollback(err, tx, "Error Get Transaction")
	if err != nil {
		return nil, err
	}

	if cache := service.CacheData[int(cusid)]; cache != nil {
		data, _ := cache.([]model.TransactionCus)
		if len(data) == len(items) {
			return data, nil
		}
	}
	wg := sync.WaitGroup{}
	limit := make(chan struct{}, 30)
	defer close(limit)
	for index, itemsss := range items {
		wg.Add(1)
		limit <- struct{}{}
		go func(itemss model.TransactionCus, index int64) {
			defer func() { <-limit }()
			defer wg.Done()
			tx2 := txpool.Get().(*sql.Tx)
			defer tx2.Commit()
			ctx2 := context.Background()
			paymentInfo, _ := service.PaymentsRepository.GetAllPaymentByid(ctx2, tx2, itemss.PaymentId)
			items[index].PaymentInfo = paymentInfo
			orderitemss := service.TransactionRepository.GetAllOrderItems(ctx2, tx2, itemss.TransactionId)
			items[index].OrderItem = orderitemss
		}(itemsss, int64(index))
	}
	wg.Wait()
	if service.CacheData[int(cusid)] == nil || service.CacheData[int(cusid)] != nil {
		if service.CacheData[int(cusid)] != nil {
			cache := service.CacheData[int(cusid)].([]model.TransactionCus)
			if len(cache) != len(items) {
				service.CacheData[int(cusid)] = items
				return items, nil
			}
		}
		service.CacheData[int(cusid)] = items
	}
	return items, nil
}
func (service *InitService) InsertTmpTransaction(ctx context.Context, req model.TempTransactionRequest, csid int64) error {
	err := service.Validate.Struct(req)
	if err != nil {
		return err
	}
	tx, _ := service.DB.Begin()
	qty, err2 := service.TransactionRepository.CheckIfExisttmp(context.Background(), tx, req.ProductId, csid)
	tx2, _ := service.DB.Begin()
	if err2 != nil {
		req.Qty = int64(qty) + req.Qty
		err3 := service.TransactionRepository.InsertTempTransaction(ctx, tx2, req, csid)
		return helper.TxRollback(err3, tx2, "nil")
	}
	var payload = model.TempUpdateTransactionRequest{
		Productid:  req.ProductId,
		Qty:        req.Qty + int64(qty),
		Customerid: csid,
	}
	err3 := service.TransactionRepository.UpdateTempTransaction(ctx, tx2, payload)
	return helper.TxRollback(err3, tx2, "nil")

}

func (service *InitService) FindAllTrxByTransactionid(ctx context.Context, trxid int64, cusid int64) (*model.TransactionCus, error) {

	if cusid == -1 {
		return nil, errors.New("No Cookie Id Found")
	}
	tx, _ := service.DB.Begin()
	items, err := service.TransactionRepository.GetTransactionsByTransactionid(ctx, tx, trxid, cusid)
	if err != nil {
		return nil, err
	}
	return items, nil

}

func (service *InitService) FindTrxByTransactionid(ctx context.Context, trxid int64) (*model.TransactionAdmin, error) {

	if trxid == -1 {
		return nil, errors.New("Not Attaced Trx Id")
	}
	tx, _ := service.DB.Begin()
	defer tx.Commit()
	items, err := service.TransactionRepository.GetTransactionByTrxid(ctx, tx, trxid)
	items.OrderItem = service.TransactionRepository.GetAllOrderItems(ctx, tx, trxid)
	if err != nil {
		return nil, err
	}
	return items, nil

}

func (service *InitService) UploadProof(ctx context.Context, proof string, trxid string) error {
	if trxid == "" {
		return errors.New("Transaction Id Not Attached")
	}
	newtrxid, _ := strconv.Atoi(trxid)

	tx, _ := service.DB.Begin()
	err := service.TransactionRepository.UpdateTransaction(ctx, tx, proof, int64(newtrxid))
	return helper.TxRollback(err, tx, "Error Update Transaction")
}
