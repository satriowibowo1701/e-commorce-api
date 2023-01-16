package controller

import (
	"net/http"

	"github.com/satriowibowo1701/e-commorce-api/config"
	"github.com/satriowibowo1701/e-commorce-api/helper"
	"github.com/satriowibowo1701/e-commorce-api/model"
)

func (controller *InitController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	config.Mutex.Lock()
	defer config.Mutex.Unlock()
	produk := model.ProdukRequest{}
	helper.ReadFromRequestBody(r, &produk)
	err := controller.ProductService.Create(r.Context(), produk)
	response := helper.ResponseWithMessage(err, "Succes Creating Product")

	helper.WriteToResponseBody(w, response, err, http.StatusBadRequest)
}

func (controller *InitController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	config.Mutex.Lock()
	defer config.Mutex.Unlock()
	produk := model.ProdukUpdate{}
	helper.ReadFromRequestBody(r, &produk)
	err := controller.ProductService.Update(r.Context(), produk)
	response := helper.ResponseWithMessage(err, "Success Updating Product")
	helper.WriteToResponseBody(w, response, err, http.StatusBadRequest)
}

func (controller *InitController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	config.Mutex.Lock()
	defer config.Mutex.Unlock()
	id := helper.GetParam("id", r)
	err := controller.ProductService.Delete(r.Context(), id)
	response := helper.ResponseWithMessage(err, "Succes Deleting Product")

	helper.WriteToResponseBody(w, response, err, http.StatusBadRequest)
}

func (controller *InitController) FindAll(w http.ResponseWriter, r *http.Request) {
	data, err := controller.ProductService.FindAll(r.Context())
	response := helper.ResponseWithData(err, data)
	helper.WriteToResponseBody(w, response, err, http.StatusBadRequest)
}
func (controller *InitController) FindAllPrdkAdmin(w http.ResponseWriter, r *http.Request) {
	data, err := controller.ProductService.FindAllPrdkAdmin(r.Context())
	response := helper.ResponseWithData(err, data)
	helper.WriteToResponseBody(w, response, err, http.StatusBadRequest)
}

func (controller *InitController) FindById(w http.ResponseWriter, r *http.Request) {
	id := helper.GetParam("id", r)
	data, err := controller.ProductService.FindById(r.Context(), id)
	response := helper.ResponseWithData(err, data)
	helper.WriteToResponseBody(w, response, err, http.StatusBadRequest)
}
