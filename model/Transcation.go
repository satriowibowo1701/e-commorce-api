package model

type TransactionAdmin struct {
	TransactionId    int64       `json:"transaction_id"`
	CustomerName     string      `json:"customer_name"`
	CustomerId       int64       `json:"customer_id"`
	Date             string      `json:"date"`
	Status           int         `json:"status"`
	TransactionTotal int64       `json:"transaction_total"`
	Destination      string      `json:"destination"`
	Proof            string      `json:"proof"`
	Payment          string      `json:"payment"`
	OrderItem        []OrderItem `json:"order_items"`
}
type TransactionCus struct {
	TransactionId    int64            `json:"transaction_id"`
	CustomerId       int64            `json:"customer_id"`
	Date             string           `json:"date"`
	Status           int              `json:"status"`
	TransactionTotal int64            `json:"transaction_total"`
	PaymentId        int64            `json:"-"`
	Destination      string           `json:"destination"`
	Proof            string           `json:"proof"`
	PaymentInfo      *PaymentResponse `json:"payment_info"`
	OrderItem        []OrderItem      `json:"order_items"`
}

type TransactionRequest struct {
	CustomerId int64        `json:"customerid"`
	Status     int64        `validate:"required" json:"status"`
	PaymentId  int64        `validate:"required" json:"payment_id"`
	Dest       string       `validate:"required" json:"destination"`
	Total      int64        `json:"total_transaction" `
	OrderItems []*OrderItem `validate:"required" json:"orderitems"`
}

type OrderItem struct {
	Id          int64  `json:"id"`
	OrderId     int64  `json:"orderid"`
	ProductId   int64  `validate:"required" json:"productid"`
	OrderQty    int64  `validate:"required" json:"orderqty"`
	OrderPrice  int64  `validate:"required" json:"orderprice"`
	ProductName string `validate:"required" json:"productname"`
}

type OrderItems struct {
	Id        int64 `json:"id"`
	ProductId int64 `json:"product_id"`
	Qty       int64 `json:"qty"`
	Price     int64 `json:"price"`
}

type TempTransactionRequest struct {
	ProductId int64 `validate:"required" json:"product_id"`
	Qty       int64 `validate:"required" json:"qty"`
	Price     int64 `validate:"required" json:"price"`
}

type TempTransaction struct {
	Id          int64  `json:"id"`
	ProductId   int64  `json:"product_id"`
	Qty         int64  `json:"qty"`
	Price       int64  `json:"price"`
	CustomerId  int64  `json:"customer_id"`
	ProductName string `json:"product_name"`
}

type TempUpdateTransactionRequest struct {
	Id         int64 `validate:"required" json:"id"`
	Qty        int64 `validate:"required" json:"qty"`
	Productid  int64
	Customerid int64
}
