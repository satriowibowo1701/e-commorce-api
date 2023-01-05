package model

type ProdukRequest struct {
	Name string `validate:"required" json:"name"`
	Qty  int64  `validate:"required" json:"qty"`
}
type Produk struct {
	Product_id int64  `json:"product_id"`
	Name       string `json:"name"`
	Qty        int64  `json:"qty"`
}
type ProdukUpdate struct {
	Product_id int64  `validate:"required" json:"product_id"`
	Name       string `validate:"required" json:"name"`
	Qty        int64  `validate:"required" json:"qty"`
}
