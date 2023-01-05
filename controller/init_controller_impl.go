package controller

import "github.com/satriowibowo1701/e-commorce-api/service"

type InitController struct {
	ProductService     service.ProductService
	UserService        service.UserService_impl
	TransactionService service.TransactionService
}

func NewInitControler(ProductService service.ProductService, TransactionService service.TransactionService, UserService service.UserService_impl) (ProdukController,
	TransactionController,
	UserController) {
	return &InitController{
			ProductService:     ProductService,
			UserService:        UserService,
			TransactionService: TransactionService,
		}, &InitController{
			ProductService:     ProductService,
			UserService:        UserService,
			TransactionService: TransactionService,
		},
		&InitController{
			ProductService:     ProductService,
			UserService:        UserService,
			TransactionService: TransactionService,
		}
}
