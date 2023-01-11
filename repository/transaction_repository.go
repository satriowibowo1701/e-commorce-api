package repository

import (
	"context"
	"database/sql"

	"github.com/satriowibowo1701/e-commorce-api/model"
)

type TransactionRepo interface {
	CreateTransaction(ctx context.Context, tx *sql.Tx, trx model.TransactionRequest) (int, error)
	GetAllTransaction(ctx context.Context, tx *sql.Tx) ([]*model.TransactionAdmin, error)
	UpdateTransaction(ctx context.Context, tx *sql.Tx, proof string, idtrx int64) error
	GetAllTransactionById(ctx context.Context, tx *sql.Tx, cstid int64) ([]*model.TransactionCus, error)
	GetTransactionByTrxid(ctx context.Context, tx *sql.Tx, trxid int64) (*model.TransactionAdmin, error)
	GetAllTransactionsByStatusCus(ctx context.Context, tx *sql.Tx, status int64, cusid int64) ([]*model.TransactionCus, error)
	InsertTempTransaction(ctx context.Context, tx *sql.Tx, trx model.TempTransactionRequest, csid int64) error
	DeleteTempTransaction(ctx context.Context, tx *sql.Tx, cid int64) error
	DeleteTempTransactionByid(ctx context.Context, tx *sql.Tx, trxid int64, cusid int64) error
	UpdateTempTransaction(ctx context.Context, tx *sql.Tx, trx model.TempUpdateTransactionRequest) error
	GetAllTempTransactionsCus(ctx context.Context, tx *sql.Tx, cusid int64) ([]*model.TempTransaction, error)
	GetTransactionsByTransactionid(ctx context.Context, tx *sql.Tx, trxid int64, cusid int64) (*model.TransactionCus, error)
	InsertOrderItems(ctx context.Context, tx *sql.Tx, trx *model.OrderItem, trxid int64) error
	GetAllOrderItems(ctx context.Context, tx *sql.Tx, transactionid int64) []model.OrderItem
	GetTempTransactionsByid(ctx context.Context, tx *sql.Tx, tmpid int64) error
	CheckIfExisttmp(ctx context.Context, tx *sql.Tx, prdctid int64, customerid int64) (int, error)
}
