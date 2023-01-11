package controller

import "net/http"

type TransactionController interface {
	CreateTransaction(w http.ResponseWriter, r *http.Request)
	GetAllTempTransactionsCus(w http.ResponseWriter, r *http.Request)
	GetAllTransactionsCus(w http.ResponseWriter, r *http.Request)
	GetAllTransactionsByStatusCus(w http.ResponseWriter, r *http.Request)
	GetAllTransactionsByIdCus(w http.ResponseWriter, r *http.Request)
	InsertTransactionsTmp(w http.ResponseWriter, r *http.Request)
	DeleteTransactionsTmp(w http.ResponseWriter, r *http.Request)
	UpdateTransactionsTmp(w http.ResponseWriter, r *http.Request)
	GetTransactionsCusByTrxid(w http.ResponseWriter, r *http.Request)
	Image(w http.ResponseWriter, r *http.Request)
	Upload(w http.ResponseWriter, r *http.Request)
}
