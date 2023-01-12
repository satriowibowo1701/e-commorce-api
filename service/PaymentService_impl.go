package service

import (
	"context"
	"errors"

	"github.com/satriowibowo1701/e-commorce-api/helper"
	"github.com/satriowibowo1701/e-commorce-api/model"
)

func (service *InitService) CreatePayment(ctx context.Context, req *model.PaymentRequest) error {

	err := service.Validate.Struct(req)
	if err != nil {
		return err
	}
	tx, _ := service.DB.Begin()
	_, err2 := service.PaymentsRepository.GetAllPaymentBynumber(ctx, tx, req.CardNum)
	if err2 == nil {
		return helper.TxRollback(errors.New(""), tx, "CardNumber Already Exists")
	}

	err1 := service.PaymentsRepository.CreatePayment(ctx, tx, req)

	return helper.TxRollback(err1, tx, "Error Create Payment")
}

func (service *InitService) UpdatePayment(ctx context.Context, req *model.UpdatePaymentRequest) error {

	err := service.Validate.Struct(req)
	if err != nil {
		return err
	}
	tx, _ := service.DB.Begin()
	err2 := service.PaymentsRepository.CheckifExist(ctx, tx, req.CardNum, req.CardName)
	if err2 == nil {
		return helper.TxRollback(errors.New("error"), tx, "CardNumber Already Exists")
	}
	err3 := service.PaymentsRepository.UpdatePayment(ctx, tx, req)

	return helper.TxRollback(err3, tx, "Error Create Payment")
}

func (service *InitService) DeletePayment(ctx context.Context, id int64) error {
	tx, _ := service.DB.Begin()
	err := service.PaymentsRepository.DeletePayment(ctx, tx, id)
	return helper.TxRollback(err, tx, "Cannot Delete Payment")
}

func (service *InitService) GetAllPayments(ctx context.Context) ([]*model.PaymentResponse, error) {
	tx, _ := service.DB.Begin()
	payments, err := service.PaymentsRepository.GetAllPayment(ctx, tx)
	defer helper.TxRollback(err, tx, "Cannnot Get Payment")
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func (service *InitService) GetPaymentByid(ctx context.Context, id int64) (*model.PaymentResponse, error) {
	tx, _ := service.DB.Begin()
	payment, err := service.PaymentsRepository.GetAllPaymentByid(ctx, tx, id)
	defer helper.TxRollback(err, tx, "Cannnot Get Payment")
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (service *InitService) GetPaymentByName(ctx context.Context, name string) ([]*model.PaymentResponse, error) {
	tx, _ := service.DB.Begin()
	if name == "" {
		return nil, errors.New("Param Name Not Found")
	}
	payment, err := service.PaymentsRepository.GetAllPaymentByholdername(ctx, tx, name)
	defer helper.TxRollback(err, tx, "Cannnot Get Payment")
	if err != nil {
		return nil, err
	}
	return payment, nil
}
