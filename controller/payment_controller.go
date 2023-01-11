package controller

import "net/http"

type PaymentController interface {
	CreatePayment(w http.ResponseWriter, r *http.Request)
	UpdatePayment(w http.ResponseWriter, r *http.Request)
	DeletePayment(w http.ResponseWriter, r *http.Request)
	FindAllPayment(w http.ResponseWriter, r *http.Request)
	FindPaymentById(w http.ResponseWriter, r *http.Request)
	FindPaymentByName(w http.ResponseWriter, r *http.Request)
}
