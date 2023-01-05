package model

type WebResponseWithData struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type WebResponseWithMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
