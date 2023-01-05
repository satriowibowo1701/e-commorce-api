package model

type TransactionAdmin struct {
	TransactionId int64       `json:"transaction_id"`
	CustomerName  string      `json:"customer_name"`
	CustomerId    int64       `json:"customer_id"`
	Date          string      `json:"date"`
	Status        int         `json:"status"`
	OrderItem     []OrderItem `json:"order_items"`
}
type TransactionCus struct {
	TransactionId int64       `json:"transaction_id"`
	CustomerId    int64       `json:"customer_id"`
	Date          string      `json:"date"`
	Status        int         `json:"status"`
	OrderItem     []OrderItem `json:"order_items"`
}

type TransactionRequest struct {
	TransactionId int64        `json:"transaction_id"`
	CustomerId    int64        `validate:"required" json:"customerid"`
	Status        int64        `validate:"required" json:"status"`
	OrderItems    []*OrderItem `validate:"required" json:"orderitems"`
}

type OrderItem struct {
	Id          int64  `json:"id"`
	OrderId     int64  `json:"orderid"`
	ProductId   int64  `json:"productid"`
	OrderQty    int64  `json:"orderqty"`
	OrderPrice  int64  `json:"orderprice"`
	ProductName string `json:"productname"`
}

type OrderItems struct {
	Id        int64 `json:"id"`
	ProductId int64 `json:"product_id"`
	Qty       int64 `json:"qty"`
	Price     int64 `json:"price"`
}

type TempTransactionRequest struct {
	ProductId  int64 `validate:"required" json:"product_id"`
	Qty        int64 `validate:"required" json:"qty"`
	Price      int64 `validate:"required" json:"price"`
	CustomerId int64 `validate:"required" json:"customer_id"`
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
	Id  int64 `validate:"required" json:"id"`
	Qty int64 `validate:"required" json:"qty"`
}
