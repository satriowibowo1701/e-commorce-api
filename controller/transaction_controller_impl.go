package controller

import (
	"log"
	"net/http"

	"github.com/satriowibowo1701/e-commorce-api/config"
	"github.com/satriowibowo1701/e-commorce-api/helper"
	"github.com/satriowibowo1701/e-commorce-api/model"
)

func (controller *InitController) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	config.Mutex.Lock()
	defer config.Mutex.Unlock()
	transaction := model.TransactionRequest{}
	helper.ReadFromRequestBody(r, &transaction)
	log.Println(transaction)
	err := controller.TransactionService.CreateTransaction(r.Context(), transaction)
	response := helper.ResponseWithMessage(err, "Succes Creating Transaction")
	helper.WriteToResponseBody(w, response)
}

func (controller *InitController) GetAllTransactionsCus(w http.ResponseWriter, r *http.Request) {
	data, err := controller.TransactionService.FindAllTransactionCustomer(r.Context())

	response := helper.ResponseWithData(err, data)

	helper.WriteToResponseBody(w, response)
}

func (controller *InitController) GetAllTransactionsByStatusCus(w http.ResponseWriter, r *http.Request) {
	cusid := model.User{}
	helper.ReadFromRequestBody(r, &cusid)
	status := helper.GetParam("status", r)
	data, err := controller.TransactionService.FindAllTransactionByStatus(r.Context(), status, cusid.ID)
	response := helper.ResponseWithData(err, data)
	helper.WriteToResponseBody(w, response)
}

func (controller *InitController) GetAllTransactionsByIdCus(w http.ResponseWriter, r *http.Request) {
	cusid := model.User{}
	helper.ReadFromRequestBody(r, &cusid)
	data, err := controller.TransactionService.FindAllTransactionById(r.Context(), cusid.ID)
	response := helper.ResponseWithData(err, data)
	helper.WriteToResponseBody(w, response)
}

func (controller *InitController) InsertTransactionsTmp(w http.ResponseWriter, r *http.Request) {
	config.Mutex.Lock()
	defer config.Mutex.Unlock()
	data := model.TempTransactionRequest{}
	helper.ReadFromRequestBody(r, &data)
	err := controller.TransactionService.InsertTmpTransaction(r.Context(), data)
	response := helper.ResponseWithMessage(err, "Successfully insert tmp transaction")
	helper.WriteToResponseBody(w, response)
}

func (controller *InitController) DeleteTransactionsTmp(w http.ResponseWriter, r *http.Request) {
	config.Mutex.Lock()
	defer config.Mutex.Unlock()
	id := helper.GetParam("id", r)
	err := controller.TransactionService.DeleteTmpTransaction(r.Context(), int64(id))
	response := helper.ResponseWithMessage(err, "Successfully delete tmp transaction")
	helper.WriteToResponseBody(w, response)
}
func (controller *InitController) UpdateTransactionsTmp(w http.ResponseWriter, r *http.Request) {
	config.Mutex.Lock()
	defer config.Mutex.Unlock()
	data := model.TempUpdateTransactionRequest{}
	helper.ReadFromRequestBody(r, &data)
	err := controller.TransactionService.UpdateTmpTransaction(r.Context(), data)
	response := helper.ResponseWithMessage(err, "Successfully Update tmp transaction")
	helper.WriteToResponseBody(w, response)
}

func (controller *InitController) GetAllTempTransactionsCus(w http.ResponseWriter, r *http.Request) {
	config.Mutex.Lock()
	defer config.Mutex.Unlock()
	cusdata := model.User{}
	helper.ReadFromRequestBody(r, &cusdata)
	data, err := controller.TransactionService.FindAllTmpTransactionCustomer(r.Context(), int64(cusdata.ID))
	response := helper.ResponseWithData(err, data)
	helper.WriteToResponseBody(w, response)
}
