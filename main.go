package main

import (
	"math/rand"
	"net/http"
	"runtime"
	"time"

	"github.com/go-playground/validator"
	"github.com/satriowibowo1701/e-commorce-api/config"
	"github.com/satriowibowo1701/e-commorce-api/controller"
	"github.com/satriowibowo1701/e-commorce-api/db"
	"github.com/satriowibowo1701/e-commorce-api/helper"
	"github.com/satriowibowo1701/e-commorce-api/middleware"
	"github.com/satriowibowo1701/e-commorce-api/repository"
	"github.com/satriowibowo1701/e-commorce-api/router"
	"github.com/satriowibowo1701/e-commorce-api/service"
)

func main() {

	runtime.GOMAXPROCS(4)
	err := db.Newmigrate()
	rand.Seed(time.Now().UnixNano())
	helper.PanicIfError(err)
	db := db.NewDB()
	defer db.Close()
	validate := validator.New()
	productrepo := repository.NewProductRepo()
	Transactionrepo := repository.NewTransactionRepository()
	UserRepo := repository.NewUserRepository()
	PaymentRepo := repository.NewPaymentRepo()
	router := router.NewRouter(controller.NewInitControler(service.RunService(db, validate, UserRepo, productrepo, Transactionrepo, PaymentRepo)))
	server := http.Server{
		Addr:    "localhost:" + config.PORT,
		Handler: middleware.AuthtenticationMiddleware(router),
	}
	server.ListenAndServe()

}
