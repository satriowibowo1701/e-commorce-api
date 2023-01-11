package helper

import (
	"database/sql"
	"errors"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func TxRollback(err error, tx *sql.Tx, message string) error {
	if err != nil {
		tx.Rollback()
		return errors.New(message)
	}
	tx.Commit()
	return nil
}

func TxRollTrx(err error, tx *sql.Tx, message string, data int) (error, int) {
	if err != nil {
		tx.Rollback()
		return errors.New(message), 0
	}
	tx.Commit()
	return nil, data
}

func IfError(err error, message string) error {
	if err != nil {
		return errors.New(message)
	}
	return nil
}
