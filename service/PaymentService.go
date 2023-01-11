package service

import (
	"context"

	"github.com/satriowibowo1701/e-commorce-api/model"
)

type PaymentService interface {
	CreatePayment(ctx context.Context, req *model.PaymentRequest) error
	UpdatePayment(ctx context.Context, req *model.UpdatePaymentRequest) error
	DeletePayment(ctx context.Context, id int64) error
	GetAllPayments(ctx context.Context) ([]*model.PaymentResponse, error)
	GetPaymentByid(ctx context.Context, id int64) (*model.PaymentResponse, error)
	GetPaymentByName(ctx context.Context, name string) ([]*model.PaymentResponse, error)
}
