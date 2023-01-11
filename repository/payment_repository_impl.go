package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/satriowibowo1701/e-commorce-api/helper"
	"github.com/satriowibowo1701/e-commorce-api/model"
)

type Paymentimpl struct {
}

func NewPaymentRepo() PaymentRepo {
	return &Paymentimpl{}
}

func (p *Paymentimpl) CreatePayment(ctx context.Context, tx *sql.Tx, req *model.PaymentRequest) error {
	SQL := "insert into payments(name,norek,cardholdername) values ($1,$2,$3)"
	_, err := tx.ExecContext(ctx, SQL, req.CardName, req.CardNum, req.CardHolderName)
	return helper.IfError(err, "error Creating Payment")
}

func (p *Paymentimpl) DeletePayment(ctx context.Context, tx *sql.Tx, id int64) error {
	SQL := "DELETE from payments WHERE id=$1"
	_, err := tx.ExecContext(ctx, SQL, id)
	return helper.IfError(err, "error Delete Payment")
}

func (p *Paymentimpl) UpdatePayment(ctx context.Context, tx *sql.Tx, req *model.UpdatePaymentRequest) error {
	SQL := "UPDATE payments SET name=$1,norek=$2,cardholdername=$3 WHERE id=$4"
	_, err := tx.ExecContext(ctx, SQL, req.CardName, req.CardNum, req.CardHolderName, req.Id)
	return helper.IfError(err, "error Updating Produk")
}

func (p *Paymentimpl) GetAllPayment(ctx context.Context, tx *sql.Tx) ([]*model.PaymentResponse, error) {
	SQL := "SELECT id,name,norek,cardholdername FROM payments"
	rows, err := tx.QueryContext(ctx, SQL)
	payments := []*model.PaymentResponse{}
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		payment := &model.PaymentResponse{}
		err = rows.Scan(&payment.Id, &payment.CardName, &payment.CardNum, &payment.CardHolderName)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}
	return payments, nil
}

func (p *Paymentimpl) GetAllPaymentByid(ctx context.Context, tx *sql.Tx, id int64) (*model.PaymentResponse, error) {
	SQL := "SELECT id,name,norek,cardholdername FROM payments WHERE id=$1"
	rows, err := tx.QueryContext(ctx, SQL, id)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		payment := &model.PaymentResponse{}
		err := rows.Scan(&payment.Id, &payment.CardName, &payment.CardNum, &payment.CardHolderName)
		if err != nil {

			return nil, err
		}
		return payment, nil
	}
	return nil, errors.New("No Data Found")
}

func (p *Paymentimpl) GetAllPaymentByholdername(ctx context.Context, tx *sql.Tx, name string) ([]*model.PaymentResponse, error) {
	SQL := "SELECT id,name,norek,cardholdername FROM payments WHERE cardholdername=$1"
	rows, err := tx.QueryContext(ctx, SQL, name)
	if err != nil {
		return nil, err

	}
	payments := []*model.PaymentResponse{}
	for rows.Next() {
		payment := &model.PaymentResponse{}
		err = rows.Scan(&payment.Id, &payment.CardName, &payment.CardNum, &payment.CardHolderName)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}
	return payments, nil
}

func (p *Paymentimpl) GetAllPaymentBynumber(ctx context.Context, tx *sql.Tx, number int64) (*model.PaymentResponse, error) {
	SQL := "SELECT id,name,norek,cardholdername FROM payments WHERE norek=$1"
	rows, err := tx.QueryContext(ctx, SQL, number)

	if rows.Next() {
		payment := &model.PaymentResponse{}
		err = rows.Scan(&payment.Id, &payment.CardName, &payment.CardNum, &payment.CardHolderName)

		if err != nil {
			return nil, err
		}
		return payment, nil
	}
	return nil, errors.New("No Data Found")
}

func (p *Paymentimpl) CheckifExist(ctx context.Context, tx *sql.Tx, number int64, name string) error {
	SQL := "SELECT id FROM payments WHERE norek=$1 and name=$2"
	row, err := tx.QueryContext(ctx, SQL, number, name)
	if err != nil {
		return err
	}
	if row.Next() {
		return nil
	}
	return errors.New("No Data Found")
}
