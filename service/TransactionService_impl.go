package service

import (
	"context"
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
	var newprice int64
	for _, items := range request.OrderItems {
		newprice += items.OrderPrice * items.OrderQty
	}
	request.Total = newprice
	tx, _ := service.DB.Begin()
	request.CustomerId = csid
	id, err2 := service.TransactionRepository.CreateTransaction(ctx, tx, request)
	defer helper.TxRollback(err2, tx, "Error Create Trx")
	if err2 != nil {
		return err2
	}
	go service.TransactionRepository.DeleteTempTransaction(ctx, tx, request.CustomerId)
	wg := sync.WaitGroup{}
	for _, items := range request.OrderItems {
		wg.Add(2)
		go func(itemss *model.OrderItem) {
			defer wg.Done()
			tx3, _ := service.DB.Begin()
			defer tx3.Commit()
			service.TransactionRepository.InsertOrderItems(context.Background(), tx3, itemss, int64(id))
		}(items)
		go func(itemss *model.OrderItem) {
			ctx2 := context.Background()
			tx2, _ := service.DB.Begin()
			defer wg.Done()
			defer tx2.Commit()
			data, _ := service.ProdukRepostory.FindById(ctx2, tx2, int(itemss.ProductId))
			service.ProdukRepostory.UpdateQty(ctx2, tx2, data.Qty-itemss.OrderQty, itemss.ProductId)
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

func (service *InitService) FindAllTransactionCustomer(ctx context.Context) ([]*model.TransactionAdmin, error) {
	tx, _ := service.DB.Begin()
	items, err := service.TransactionRepository.GetAllTransaction(ctx, tx)
	if err != nil {
		return nil, err
	}
	tx.Commit()
	wg := sync.WaitGroup{}
	for _, itemsss := range items {
		wg.Add(1)
		go func(itemss *model.TransactionAdmin) {
			defer wg.Done()
			tx2, _ := service.DB.Begin()
			defer tx2.Commit()
			itemss.OrderItem = service.TransactionRepository.GetAllOrderItems(context.Background(), tx2, itemss.TransactionId)
		}(itemsss)
	}
	wg.Wait()
	return items, nil
}

func (service *InitService) FindAllTransactionByStatus(ctx context.Context, status int, cusid int) ([]*model.TransactionCus, error) {
	if cusid == -1 {
		return nil, errors.New("No Cookie Id Found")
	}
	tx, _ := service.DB.Begin()
	items, err := service.TransactionRepository.GetAllTransactionsByStatusCus(ctx, tx, int64(status), int64(cusid))
	defer helper.TxRollback(err, tx, "Error Get Transaction")
	if err != nil {
		return nil, err
	}
	wg := sync.WaitGroup{}
	for _, itemsss := range items {
		wg.Add(2)
		go func(itemss *model.TransactionCus) {
			defer wg.Done()
			tx3, _ := service.DB.Begin()
			defer tx3.Commit()

			orderitems := service.TransactionRepository.GetAllOrderItems(context.Background(), tx3, itemss.TransactionId)
			itemss.OrderItem = orderitems

		}(itemsss)
		go func(itemss *model.TransactionCus) {
			defer wg.Done()
			tx2, _ := service.DB.Begin()
			defer tx2.Commit()
			paymentInfo, _ := service.PaymentsRepository.GetAllPaymentByid(context.Background(), tx2, itemss.PaymentId)
			itemss.PaymentInfo = paymentInfo
		}(itemsss)
	}
	wg.Wait()

	return items, nil
}

func (service *InitService) FindAllTransactionById(ctx context.Context, cusid int) ([]*model.TransactionCus, error) {
	if cusid == -1 {
		return nil, errors.New("No Cookie Id Found")
	}
	tx, _ := service.DB.Begin()
	items, err := service.TransactionRepository.GetAllTransactionById(ctx, tx, int64(cusid))

	helper.TxRollback(err, tx, "Error Get Transaction")
	if err != nil {
		return nil, err
	}
	wg := sync.WaitGroup{}
	for _, itemsss := range items {
		wg.Add(2)
		go func(itemss *model.TransactionCus) {
			defer wg.Done()
			tx3, _ := service.DB.Begin()
			defer tx3.Commit()
			orderitemss := service.TransactionRepository.GetAllOrderItems(context.Background(), tx3, itemss.TransactionId)
			itemss.OrderItem = orderitemss
		}(itemsss)
		go func(itemss *model.TransactionCus) {
			defer wg.Done()
			ctx2 := context.Background()
			tx2, _ := service.DB.Begin()
			defer tx2.Commit()
			paymentInfo, _ := service.PaymentsRepository.GetAllPaymentByid(ctx2, tx2, itemss.PaymentId)
			itemss.PaymentInfo = paymentInfo
		}(itemsss)
	}
	wg.Wait()

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
