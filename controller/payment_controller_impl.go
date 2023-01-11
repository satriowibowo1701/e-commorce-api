package controller

import (
	"net/http"

	"github.com/satriowibowo1701/e-commorce-api/config"
	"github.com/satriowibowo1701/e-commorce-api/helper"
	"github.com/satriowibowo1701/e-commorce-api/model"
)

func (controller *InitController) CreatePayment(w http.ResponseWriter, r *http.Request) {
	config.Mutex.Lock()
	defer config.Mutex.Unlock()
	payment := model.PaymentRequest{}
	helper.ReadFromRequestBody(r, &payment)
	err := controller.PaymentService.CreatePayment(r.Context(), &payment)
	response := helper.ResponseWithMessage(err, "Success Creating Payment")
	helper.WriteToResponseBody(w, response, err, http.StatusBadRequest)
}

func (controller *InitController) UpdatePayment(w http.ResponseWriter, r *http.Request) {
	payment := &model.UpdatePaymentRequest{}
	helper.ReadFromRequestBody(r, &payment)
	err := controller.PaymentService.UpdatePayment(r.Context(), payment)
	response := helper.ResponseWithMessage(err, "Success Updating Payment")
	helper.WriteToResponseBody(w, response, err, http.StatusBadRequest)
}

func (controller *InitController) DeletePayment(w http.ResponseWriter, r *http.Request) {
	config.Mutex.Lock()
	defer config.Mutex.Unlock()
	id := helper.GetParam("id", r)
	err := controller.PaymentService.DeletePayment(r.Context(), int64(id))
	response := helper.ResponseWithMessage(err, "Success Deleting Payment")

	helper.WriteToResponseBody(w, response, err, http.StatusBadRequest)
}

func (controller *InitController) FindAllPayment(w http.ResponseWriter, r *http.Request) {
	data, err := controller.PaymentService.GetAllPayments(r.Context())
	response := helper.ResponseWithData(err, data)
	helper.WriteToResponseBody(w, response, err, http.StatusBadRequest)
}

func (controller *InitController) FindPaymentById(w http.ResponseWriter, r *http.Request) {
	id := helper.GetParam("id", r)
	data, err := controller.PaymentService.GetPaymentByid(r.Context(), int64(id))
	response := helper.ResponseWithData(err, data)
	helper.WriteToResponseBody(w, response, err, http.StatusBadRequest)
}
func (controller *InitController) FindPaymentByName(w http.ResponseWriter, r *http.Request) {
	name := helper.GetParamString("name", r)
	data, err := controller.PaymentService.GetPaymentByName(r.Context(), name)
	response := helper.ResponseWithData(err, data)
	helper.WriteToResponseBody(w, response, err, http.StatusBadRequest)
}
