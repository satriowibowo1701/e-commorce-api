package repository

import (
	"context"
	"database/sql"

	"github.com/satriowibowo1701/e-commorce-api/model"
)

type PaymentRepo interface {
	CreatePayment(ctx context.Context, tx *sql.Tx, req *model.PaymentRequest) error
	DeletePayment(ctx context.Context, tx *sql.Tx, id int64) error
	UpdatePayment(ctx context.Context, tx *sql.Tx, req *model.UpdatePaymentRequest) error
	GetAllPayment(ctx context.Context, tx *sql.Tx) ([]*model.PaymentResponse, error)
	GetAllPaymentByid(ctx context.Context, tx *sql.Tx, id int64) (*model.PaymentResponse, error)
	GetAllPaymentByholdername(ctx context.Context, tx *sql.Tx, name string) ([]*model.PaymentResponse, error)
	GetAllPaymentBynumber(ctx context.Context, tx *sql.Tx, number int64) (*model.PaymentResponse, error)
	CheckifExist(ctx context.Context, tx *sql.Tx, number int64, name string) error
}
