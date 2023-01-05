package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/satriowibowo1701/e-commorce-api/helper"
	"github.com/satriowibowo1701/e-commorce-api/model"
)

type Transaction struct {
	Db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepo {
	return &Transaction{Db: db}
}

func (t *Transaction) CreateTransaction(ctx context.Context, tx *sql.Tx, trx model.TransactionRequest) (error, int) {
	SQL := "insert into transaction(customer_id,date,status) values (?,?,?)"
	res, err := tx.ExecContext(ctx, SQL, trx.CustomerId, time.Now(), trx.Status)
	trxid, _ := res.LastInsertId()

	return helper.TxRollTrx(err, tx, "error Create Transaction", int(trxid))

}

func (t *Transaction) GetAllTransaction(ctx context.Context, tx *sql.Tx) ([]*model.TransactionAdmin, error) {
	SQL := "select t.transaction_id,u.name,t.customer_id,t.date,t.status from user u JOIN transaction t on u.id=t.customer_id"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer tx.Commit()
	defer rows.Close()
	var transactions []*model.TransactionAdmin
	for rows.Next() {
		trx := model.TransactionAdmin{}
		err := rows.Scan(&trx.TransactionId, &trx.CustomerName, &trx.CustomerId, &trx.Date, &trx.Status)
		if err != nil {
			return nil, errors.New("Cannot Scaning")
		}

		transactions = append(transactions, &trx)
	}
	return transactions, nil
}

func (t *Transaction) GetAllTransactionById(ctx context.Context, tx *sql.Tx, cstid int64) ([]*model.TransactionCus, error) {
	SQL := "select transaction_id,customer_id,date,status from transaction WHERE customer_id=?"
	rows, err := tx.QueryContext(ctx, SQL, cstid)
	if err != nil {
		return nil, errors.New("Error Sql")
	}
	defer rows.Close()
	defer tx.Commit()
	var transactions []*model.TransactionCus
	for rows.Next() {
		trx := &model.TransactionCus{}
		err := rows.Scan(&trx.TransactionId, &trx.CustomerId, &trx.Date, &trx.Status)
		if err != nil {
			return nil, errors.New("Cannot Scaning")
		}

		transactions = append(transactions, trx)
	}
	return transactions, nil
}

func (t *Transaction) GetAllTransactionsByStatusCus(ctx context.Context, tx *sql.Tx, status int64, cusid int64) ([]*model.TransactionCus, error) {
	SQL := "select transaction_id,customer_id,date,status from transaction WHERE status=? AND customer_id=?"
	fmt.Println(cusid, status)
	rows, err := tx.QueryContext(ctx, SQL, status, cusid)

	if err != nil {
		return nil, errors.New("Error Sql")
	}
	defer rows.Close()
	defer tx.Commit()
	var transactions []*model.TransactionCus
	for rows.Next() {
		trx := model.TransactionCus{}
		err := rows.Scan(&trx.TransactionId, &trx.CustomerId, &trx.Date, &trx.Status)
		if err != nil {
			return nil, errors.New("Cannot Scaning")
		}
		transactions = append(transactions, &trx)
	}
	return transactions, nil
}

func (t *Transaction) InsertTempTransaction(ctx context.Context, tx *sql.Tx, trx model.TempTransactionRequest) error {
	SQL := "insert into temp_order(product_id,qty,price,customer_id) values (?,?,?,?)"
	_, err := tx.ExecContext(ctx, SQL, trx.ProductId, trx.Qty, trx.Price, trx.CustomerId)

	return helper.TxRollback(err, tx, "Cannot Insert Tmp Transaction")
}

func (t *Transaction) UpdateTempTransaction(ctx context.Context, tx *sql.Tx, trx model.TempUpdateTransactionRequest) error {
	SQL := "UPDATE temp_order SET qty = ? WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL, trx.Qty, trx.Id)
	return helper.TxRollback(err, tx, "Cannot Update Tmp Transaction")
}
func (t *Transaction) DeleteTempTransaction(ctx context.Context, tx *sql.Tx, cid int64) error {
	SQL := "DELETE FROM temp_order WHERE customer_id = ?"
	_, err := tx.ExecContext(ctx, SQL, cid)

	return helper.TxRollback(err, tx, "Cannot Insert to Tmp Transaction")
}
func (t *Transaction) GetAllTempTransactionsCus(ctx context.Context, tx *sql.Tx, cusid int64) ([]*model.TempTransaction, error) {
	SQL := "select t.id, t.product_id, t.qty, t.price, t.customer_id, p.name from t temp_order JOIN p product ON p.product_id=t.product_id WHERE t.customer_id=?"
	rows, err := tx.QueryContext(ctx, SQL, cusid)
	if err != nil {
		return nil, errors.New("Error Sql TMP")
	}
	defer tx.Commit()
	defer rows.Close()
	var transactions []*model.TempTransaction
	for rows.Next() {
		trx := model.TempTransaction{}
		err := rows.Scan(&trx.Id, &trx.ProductId, &trx.Qty, &trx.Price, &trx.CustomerId, &trx.ProductName)
		if err != nil {
			return nil, errors.New("Cannot Scaning")
		}
		transactions = append(transactions, &trx)
	}
	return transactions, nil

}

func (t *Transaction) GetTempTransactionsByid(ctx context.Context, tx *sql.Tx, tmpid int64) error {
	SQL := "Select id from temp_order where id=?"
	rows, err := tx.QueryContext(ctx, SQL, tmpid)
	if err != nil {
		return errors.New("Error Sql TMP")
	}
	defer rows.Close()
	if rows.Next() {
		return nil
	}
	return errors.New("Transaction temp not found")

}
func (t *Transaction) GetAllOrderItems(ctx context.Context, tx *sql.Tx, transactionid int64) []model.OrderItem {

	SQL := "select t.id, t.transaction_id,t.product_id, t.qty, t.price, p.name from transaction_items t JOIN product p on t.product_id = p.product_id WHERE t.transaction_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, transactionid)
	defer rows.Close()
	defer tx.Commit()
	if err != nil {
		return nil
	}
	items := []model.OrderItem{}
	item := model.OrderItem{}
	for rows.Next() {
		rows.Scan(&item.Id, &item.OrderId, &item.ProductId, &item.OrderQty, &item.OrderPrice, &item.ProductName)
		items = append(items, item)
	}
	return items
}

func (t *Transaction) InsertOrderItems(ctx context.Context, tx *sql.Tx, trx *model.OrderItem, trxid int64) {

	SQL := "insert into transaction_items(transaction_id,product_id,qty,price) values (?,?,?,?)"
	_, err := tx.ExecContext(ctx, SQL, trxid, trx.ProductId, trx.OrderQty, trx.OrderPrice)
	helper.TxRollback(err, tx, "a")

}
