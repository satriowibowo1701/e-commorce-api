package controller

import "net/http"

type UserController interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	FindAllUsers(w http.ResponseWriter, r *http.Request)
}
