package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/satriowibowo1701/e-commorce-api/controller"
	"github.com/satriowibowo1701/e-commorce-api/db"
	"github.com/satriowibowo1701/e-commorce-api/repository"
	"github.com/satriowibowo1701/e-commorce-api/router"
	"github.com/satriowibowo1701/e-commorce-api/service"
)

func main() {
	// runtime.GOMAXPROCS(3)

	db := db.NewDB()
	defer db.Close()
	validate := validator.New()
	productrepo := repository.NewProductRepo()
	Transactionrepo := repository.NewTransactionRepository(db)
	UserRepo := repository.NewUserRepository()

	router := router.NewRouter(controller.NewInitControler(service.RunService(db, validate, UserRepo, productrepo, Transactionrepo)))

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}
	fmt.Println("Listening")
	server.ListenAndServe()

}
