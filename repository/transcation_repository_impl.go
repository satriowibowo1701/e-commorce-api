package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/satriowibowo1701/e-commorce-api/helper"
	"github.com/satriowibowo1701/e-commorce-api/model"
)

type Transaction struct {
}

func NewTransactionRepository() TransactionRepo {
	return &Transaction{}
}

func (t *Transaction) CreateTransaction(ctx context.Context, tx *sql.Tx, trx model.TransactionRequest) (int, error) {
	SQL := "INSERT INTO transaction(customer_id,date,status,total,payment_id,alamat_pengiriman,bukti_pembayaran) values($1,$2,$3,$4,$5,$6,$7) RETURNING transaction_id"
	trxid := 0
	err := tx.QueryRowContext(ctx, SQL, trx.CustomerId, time.Now(), trx.Status, trx.Total, trx.PaymentId, trx.Dest, "waiting").Scan(&trxid)
	if err != nil {
		return 0, err
	}
	return trxid, nil

}

func (t *Transaction) GetAllTransaction(ctx context.Context, tx *sql.Tx) ([]model.TransactionAdmin, error) {
	SQL := "select t.transaction_id, t.customer_id, t.date, t.status, t.total,u.name,p.name,t.alamat_pengiriman,t.bukti_pembayaran FROM usert u JOIN transaction t ON t.customer_id=u.id JOIN payments p ON p.id=t.payment_id ORDER BY t.transaction_id ASC"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var transactions []model.TransactionAdmin
	for rows.Next() {
		trx := model.TransactionAdmin{}
		err := rows.Scan(&trx.TransactionId, &trx.CustomerId, &trx.Date, &trx.Status, &trx.TransactionTotal, &trx.CustomerName, &trx.Payment, &trx.Destination, &trx.Proof)
		if err != nil {
			return nil, errors.New("Cannot Scaning")
		}
		transactions = append(transactions, trx)
	}
	return transactions, nil
}

func (t *Transaction) GetAllTransactionById(ctx context.Context, tx *sql.Tx, cstid int64) ([]model.TransactionCus, error) {
	SQL := "select transaction_id,customer_id,date,status,total,payment_id,alamat_pengiriman,bukti_pembayaran from transaction WHERE customer_id=$1 ORDER BY transaction_id ASC"
	rows, err := tx.QueryContext(ctx, SQL, cstid)
	if err != nil {
		return nil, errors.New("Error Sql")
	}
	defer rows.Close()
	var transactions []model.TransactionCus
	for rows.Next() {
		trx := model.TransactionCus{}
		err := rows.Scan(&trx.TransactionId, &trx.CustomerId, &trx.Date, &trx.Status, &trx.TransactionTotal, &trx.PaymentId, &trx.Destination, &trx.Proof)
		if err != nil {
			return nil, errors.New("Cannot Scaning")
		}

		transactions = append(transactions, trx)
	}
	return transactions, nil
}

func (t *Transaction) GetTransactionByTrxid(ctx context.Context, tx *sql.Tx, trxid int64) (*model.TransactionAdmin, error) {
	SQL := "select t.transaction_id,customer_id,u.name,t.date,t.status,t.total,p.name,t.alamat_pengiriman,t.bukti_pembayaran from transaction t JOIN usert u ON u.id=t.customer_id JOIN payments p ON p.id=t.payment_id WHERE t.transaction_id=$1 ORDER BY t.transaction_id ASC"
	rows, err := tx.QueryContext(ctx, SQL, trxid)
	if err != nil {
		return nil, errors.New("Error Sql")
	}
	defer rows.Close()
	if rows.Next() {
		trx := &model.TransactionAdmin{}
		err := rows.Scan(&trx.TransactionId, &trx.CustomerId, &trx.CustomerName, &trx.Date, &trx.Status, &trx.TransactionTotal, &trx.CustomerName, &trx.Destination, &trx.Proof)
		if err != nil {
			return nil, errors.New("Cannot Scaning")
		}

		return trx, nil
	}
	return nil, errors.New("No Data Found")
}

func (t *Transaction) GetAllTransactionsByStatusCus(ctx context.Context, tx *sql.Tx, status int64, cusid int64) ([]model.TransactionCus, error) {
	SQL := "select transaction_id,customer_id,date,status,total,payment_id,alamat_pengiriman,bukti_pembayaran from transaction WHERE status=$1 AND customer_id=$2 ORDER BY transaction_id"
	rows, err := tx.QueryContext(ctx, SQL, status, cusid)

	if err != nil {
		return nil, errors.New("Error Sql")
	}
	defer rows.Close()
	var transactions []model.TransactionCus
	for rows.Next() {
		trx := model.TransactionCus{}
		err := rows.Scan(&trx.TransactionId, &trx.CustomerId, &trx.Date, &trx.Status, &trx.TransactionTotal, &trx.PaymentId, &trx.Destination, &trx.Proof)
		if err != nil {
			return nil, errors.New("Cannot Scaning")
		}
		transactions = append(transactions, trx)
	}
	return transactions, nil
}

func (t *Transaction) GetTransactionsByTransactionid(ctx context.Context, tx *sql.Tx, trxid int64, cusid int64) (*model.TransactionCus, error) {
	SQL := "select transaction_id,customer_id,date,status,total,payment_id,alamat_pengiriman,bukti_pembayaran from transaction WHERE transaction_id=$1 AND customer_id=$2"
	rows, err := tx.QueryContext(ctx, SQL, trxid, cusid)
	if err != nil {
		return nil, errors.New("Error Sql")
	}
	defer rows.Close()
	if rows.Next() {
		trx := model.TransactionCus{}
		err := rows.Scan(&trx.TransactionId, &trx.CustomerId, &trx.Date, &trx.Status, &trx.TransactionTotal, &trx.PaymentId, &trx.Destination, &trx.Proof)
		if err != nil {
			return nil, errors.New("Cannot Scaning")
		}
		return &trx, nil
	}
	return nil, errors.New("No data found")
}

func (t *Transaction) InsertTempTransaction(ctx context.Context, tx *sql.Tx, trx model.TempTransactionRequest, csid int64) error {
	SQL := "insert into temp_order(product_id,qty,price,customer_id) values ($1,$2,$3,$4)"
	_, err := tx.ExecContext(ctx, SQL, trx.ProductId, trx.Qty, trx.Price, csid)
	return helper.IfError(err, "Cannot Insert Tmp Transaction")
}

func (t *Transaction) UpdateTempTransaction(ctx context.Context, tx *sql.Tx, trx model.TempUpdateTransactionRequest) error {
	if trx.Id != 0 {
		SQL := "UPDATE temp_order SET qty = $1 WHERE id=$2"
		_, err := tx.ExecContext(ctx, SQL, trx.Qty, trx.Id)
		return helper.IfError(err, "Cannot Update Tmp Transaction")
	} else {

		SQL := "UPDATE temp_order SET qty = $1 WHERE product_id=$2 AND customer_id=$3"
		_, err2 := tx.ExecContext(ctx, SQL, trx.Qty, trx.Productid, trx.Customerid)
		return helper.IfError(err2, "Cannot Update Tmp Transaction")
	}

}

func (t *Transaction) UpdateTransaction(ctx context.Context, tx *sql.Tx, proof string, idtrx int64) error {
	SQL := "UPDATE transaction SET bukti_pembayaran = $1 WHERE transaction_id=$2"
	_, err := tx.ExecContext(ctx, SQL, proof, idtrx)
	return helper.IfError(err, "Cannot Update Transaction")

}
func (t *Transaction) DeleteTempTransaction(ctx context.Context, tx *sql.Tx, cid int64) error {
	SQL := "DELETE FROM temp_order WHERE customer_id=$1"
	_, err := tx.ExecContext(ctx, SQL, cid)
	return helper.IfError(err, "Cannot Insert to Tmp Transaction")

}

func (t *Transaction) DeleteTempTransactionByid(ctx context.Context, tx *sql.Tx, trxid int64, cusid int64) error {
	SQL := "DELETE FROM temp_order WHERE customer_id=$1 and id=$2"
	_, err := tx.ExecContext(ctx, SQL, cusid, trxid)

	return helper.IfError(err, "Cannot Delete Tmp Transaction")
}
func (t *Transaction) GetAllTempTransactionsCus(ctx context.Context, tx *sql.Tx, cusid int64) ([]*model.TempTransaction, error) {
	SQL := "select t.id, t.product_id, t.qty, t.price, t.customer_id, p.name from temp_order t JOIN product p ON p.product_id=t.product_id WHERE t.customer_id=$1"
	rows, err := tx.QueryContext(ctx, SQL, cusid)
	if err != nil {
		return nil, errors.New("Error Sql TMP")
	}
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
	SQL := "select id from temp_order where id=$1"
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

func (t *Transaction) CheckIfExisttmp(ctx context.Context, tx *sql.Tx, prdctid int64, customerid int64) (int, error) {
	SQL := "select qty from temp_order where product_id=$1 and customer_id=$2"
	rows, err := tx.QueryContext(ctx, SQL, prdctid, customerid)
	if err != nil {
		return -1, err
	}
	defer rows.Close()
	if rows.Next() {
		qty := 0
		_ = rows.Scan(&qty)
		return qty, nil
	}
	return 0, errors.New("Not Exist")

}
func (t *Transaction) GetAllOrderItems(ctx context.Context, tx *sql.Tx, transactionid int64) []model.OrderItem {

	SQL := "select t.id, t.transaction_id,t.product_id, t.qty, t.price, p.name from transaction_items t JOIN product p on t.product_id = p.product_id WHERE t.transaction_id = $1"
	rows, err := tx.QueryContext(ctx, SQL, transactionid)
	if err != nil {
		return nil
	}
	defer rows.Close()
	items := []model.OrderItem{}
	for rows.Next() {
		item := model.OrderItem{}
		_ = rows.Scan(&item.Id, &item.OrderId, &item.ProductId, &item.OrderQty, &item.OrderPrice, &item.ProductName)
		items = append(items, item)
	}
	return items
}

func (t *Transaction) InsertOrderItems(ctx context.Context, tx *sql.Tx, trx *model.OrderItem, trxid int64) error {

	SQL := "insert into transaction_items(transaction_id,product_id,qty,price) values ($1,$2,$3,$4)"
	_, err := tx.ExecContext(ctx, SQL, trxid, trx.ProductId, trx.OrderQty, trx.OrderPrice)
	return helper.IfError(err, "Cannot Insert Tmp Transaction")

}
