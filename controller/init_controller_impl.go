package controller

import "github.com/satriowibowo1701/e-commorce-api/service"

type InitController struct {
	ProductService     service.ProductService
	UserService        service.UserService_impl
	TransactionService service.TransactionService
	PaymentService     service.PaymentService
}

func NewInitControler(ProductService service.ProductService, TransactionService service.TransactionService, UserService service.UserService_impl, PaymentService service.PaymentService) (ProdukController,
	TransactionController,
	UserController, PaymentController) {
	return &InitController{
			ProductService:     ProductService,
			UserService:        UserService,
			TransactionService: TransactionService,
			PaymentService:     PaymentService,
		}, &InitController{
			ProductService:     ProductService,
			UserService:        UserService,
			TransactionService: TransactionService,
			PaymentService:     PaymentService,
		},
		&InitController{
			ProductService:     ProductService,
			UserService:        UserService,
			TransactionService: TransactionService,
			PaymentService:     PaymentService,
		},
		&InitController{
			ProductService:     ProductService,
			UserService:        UserService,
			TransactionService: TransactionService,
			PaymentService:     PaymentService,
		}
}
