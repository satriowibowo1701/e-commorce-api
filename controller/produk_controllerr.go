package controller

import "net/http"

type ProdukController interface {
	CreateProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
	FindAllPrdkAdmin(w http.ResponseWriter, r *http.Request)
}
