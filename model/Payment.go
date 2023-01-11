package model

type PaymentRequest struct {
	CardName       string `validate:"required" json:"card_name"`
	CardNum        int64  `validate:"required" json:"card_num"`
	CardHolderName string `validate:"required" json:"card_holder_name"`
}

type UpdatePaymentRequest struct {
	Id             int    `json:"id"`
	CardName       string `validate:"required" json:"card_name"`
	CardNum        int64  `validate:"required" json:"card_num"`
	CardHolderName string `validate:"required" json:"card_holder_name"`
}

type PaymentResponse struct {
	Id             int    `json:"id"`
	CardName       string `json:"card_name"`
	CardNum        int64  `json:"card_num"`
	CardHolderName string `json:"card_holder_name"`
}
