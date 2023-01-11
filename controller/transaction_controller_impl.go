package controller

import (
	"errors"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/satriowibowo1701/e-commorce-api/config"
	"github.com/satriowibowo1701/e-commorce-api/helper"
	"github.com/satriowibowo1701/e-commorce-api/model"
)

func (controller *InitController) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	config.Mutex.Lock()
	defer config.Mutex.Unlock()
	id := helper.GetCokkie("id", r)
	transaction := model.TransactionRequest{}
	helper.ReadFromRequestBody(r, &transaction)
	err := controller.TransactionService.CreateTransaction(r.Context(), transaction, int64(id))
	response := helper.ResponseWithMessage(err, "Succes Creating Transaction")
	helper.WriteToResponseBody(w, response, err, http.StatusBadRequest)
}

func (controller *InitController) GetAllTransactionsCus(w http.ResponseWriter, r *http.Request) {
	data, err := controller.TransactionService.FindAllTransactionCustomer(r.Context())

	response := helper.ResponseWithData(err, data)

	helper.WriteToResponseBody(w, response, err, http.StatusBadRequest)
}

func (controller *InitController) GetAllTransactionsByStatusCus(w http.ResponseWriter, r *http.Request) {

	cusid := helper.GetCokkie("id", r)
	status := helper.GetParam("status", r)

	data, err := controller.TransactionService.FindAllTransactionByStatus(r.Context(), status, cusid)
	response := helper.ResponseWithData(err, data)
	helper.WriteToResponseBody(w, response, err, http.StatusBadRequest)
}

func (controller *InitController) GetAllTransactionsByIdCus(w http.ResponseWriter, r *http.Request) {

	cusid := helper.GetCokkie("id", r)
	data, err := controller.TransactionService.FindAllTransactionById(r.Context(), cusid)
	response := helper.ResponseWithData(err, data)
	helper.WriteToResponseBody(w, response, err, http.StatusBadRequest)
}

func (controller *InitController) InsertTransactionsTmp(w http.ResponseWriter, r *http.Request) {
	config.Mutex.Lock()
	defer config.Mutex.Unlock()
	cusid := make(chan int)
	go func() {
		cus := helper.GetCokkie("id", r)
		cusid <- cus
	}()
	data := model.TempTransactionRequest{}
	helper.ReadFromRequestBody(r, &data)
	err := controller.TransactionService.InsertTmpTransaction(r.Context(), data, int64(<-cusid))
	response := helper.ResponseWithMessage(err, "Successfully insert tmp transaction")
	helper.WriteToResponseBody(w, response, err, http.StatusBadRequest)
}

func (controller *InitController) DeleteTransactionsTmp(w http.ResponseWriter, r *http.Request) {
	id := helper.GetParam("id", r)
	cusid := helper.GetCokkie("id", r)
	err := controller.TransactionService.DeleteTmpTransaction(r.Context(), int64(id), int64(cusid))
	response := helper.ResponseWithMessage(err, "Successfully delete tmp transaction")
	helper.WriteToResponseBody(w, response, err, http.StatusBadRequest)
}
func (controller *InitController) UpdateTransactionsTmp(w http.ResponseWriter, r *http.Request) {
	config.Mutex.Lock()
	defer config.Mutex.Unlock()
	data := model.TempUpdateTransactionRequest{}
	helper.ReadFromRequestBody(r, &data)
	err := controller.TransactionService.UpdateTmpTransaction(r.Context(), data)
	response := helper.ResponseWithMessage(err, "Successfully Update tmp transaction")
	helper.WriteToResponseBody(w, response, err, http.StatusBadRequest)
}

func (controller *InitController) GetAllTempTransactionsCus(w http.ResponseWriter, r *http.Request) {
	cusid := helper.GetCokkie("id", r)
	data, err := controller.TransactionService.FindAllTmpTransactionCustomer(r.Context(), int64(cusid))
	response := helper.ResponseWithData(err, data)
	helper.WriteToResponseBody(w, response, err, http.StatusBadRequest)
}

func (controller *InitController) GetTransactionsCusByTrxid(w http.ResponseWriter, r *http.Request) {
	trxid := helper.GetParam("trxid", r)
	data, err := controller.TransactionService.FindTrxByTransactionid(r.Context(), int64(trxid))
	response := helper.ResponseWithData(err, data)
	helper.WriteToResponseBody(w, response, err, http.StatusBadRequest)
}

func (controller *InitController) FindTrxByTransactionId(w http.ResponseWriter, r *http.Request) {
	cusid := helper.GetCokkie("id", r)
	trxid := helper.GetParam("trxid", r)
	data, err := controller.TransactionService.FindAllTrxByTransactionid(r.Context(), int64(trxid), int64(cusid))
	response := helper.ResponseWithData(err, data)
	helper.WriteToResponseBody(w, response, err, http.StatusBadRequest)
}

func (controller *InitController) Upload(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	image, _, err := r.FormFile("proof")
	filename := "img" + strconv.Itoa(rand.Intn(int(time.Now().Unix()))) + ".jpg"
	if err != nil {
		response := helper.ResponseWithMessage(err, "No image Found")
		helper.WriteToResponseBody(w, response, err, http.StatusBadRequest)
		return
	}
	go func() {
		dest, _ := os.Create("./assets/" + filename)
		_, _ = io.Copy(dest, image)
	}()
	err2 := controller.TransactionService.UploadProof(r.Context(), filename, id)
	response := helper.ResponseWithMessage(err2, "Sukses Mengupload Bukti")
	helper.WriteToResponseBody(w, response, err2, http.StatusBadRequest)
}

func (controller *InitController) Image(w http.ResponseWriter, r *http.Request) {
	getname := helper.GetParamString("name", r)
	if getname == "" {
		response := helper.ResponseWithMessage(errors.New("File Not Exist"), "")
		helper.WriteToResponseBody(w, response, errors.New("File Not Exist"), http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, "./assets/"+getname)
}
