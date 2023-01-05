package repository

import (
	"context"
	"database/sql"

	"github.com/satriowibowo1701/e-commorce-api/model"
)

type TransactionRepo interface {
	CreateTransaction(ctx context.Context, tx *sql.Tx, trx model.TransactionRequest) (error, int)
	GetAllTransaction(ctx context.Context, tx *sql.Tx) ([]*model.TransactionAdmin, error)
	GetAllTransactionById(ctx context.Context, tx *sql.Tx, cstid int64) ([]*model.TransactionCus, error)
	GetAllTransactionsByStatusCus(ctx context.Context, tx *sql.Tx, status int64, cusid int64) ([]*model.TransactionCus, error)
	InsertTempTransaction(ctx context.Context, tx *sql.Tx, trx model.TempTransactionRequest) error
	DeleteTempTransaction(ctx context.Context, tx *sql.Tx, cid int64) error
	UpdateTempTransaction(ctx context.Context, tx *sql.Tx, trx model.TempUpdateTransactionRequest) error
	GetAllTempTransactionsCus(ctx context.Context, tx *sql.Tx, cusid int64) ([]*model.TempTransaction, error)
	InsertOrderItems(ctx context.Context, tx *sql.Tx, trx *model.OrderItem, trxid int64)
	GetAllOrderItems(ctx context.Context, tx *sql.Tx, transactionid int64) []model.OrderItem
	GetTempTransactionsByid(ctx context.Context, tx *sql.Tx, tmpid int64) error
}
