package service

import (
	"database/sql"

	"github.com/go-playground/validator"
	"github.com/satriowibowo1701/e-commorce-api/repository"
)

type InitService struct {
	DB                    *sql.DB
	Validate              *validator.Validate
	UserRepository        repository.UserRepository
	ProdukRepostory       repository.ProductRepo
	TransactionRepository repository.TransactionRepo
}

func RunService(DB *sql.DB, validate *validator.Validate, user repository.UserRepository, produk repository.ProductRepo, transaction repository.TransactionRepo) (ProductService, TransactionService, UserService_impl) {
	return &InitService{
			DB:                    DB,
			Validate:              validate,
			UserRepository:        user,
			ProdukRepostory:       produk,
			TransactionRepository: transaction,
		}, &InitService{
			DB:                    DB,
			Validate:              validate,
			UserRepository:        user,
			ProdukRepostory:       produk,
			TransactionRepository: transaction,
		}, &InitService{
			DB:                    DB,
			Validate:              validate,
			UserRepository:        user,
			ProdukRepostory:       produk,
			TransactionRepository: transaction,
		}

}
