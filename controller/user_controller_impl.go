package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/satriowibowo1701/e-commorce-api/config"
	"github.com/satriowibowo1701/e-commorce-api/helper"
	"github.com/satriowibowo1701/e-commorce-api/model"
)

func (controller *InitController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	config.Mutex.Lock()
	defer config.Mutex.Unlock()
	req := model.UserRegis{}
	helper.ReadFromRequestBody(r, &req)
	err := controller.UserService.CreateUser(r.Context(), req)
	response := helper.ResponseWithMessage(err, "Succes Creating User")
	helper.WriteToResponseBody(w, response, err, http.StatusBadRequest)
}

func (controller *InitController) LoginUser(w http.ResponseWriter, r *http.Request) {
	config.Mutex.Lock()
	defer config.Mutex.Unlock()
	req := model.LoginRequest{}
	helper.ReadFromRequestBody(r, &req)
	token, err := controller.UserService.Login(r.Context(), req)
	response := helper.ResponseWithData(err, token)
	if _, ok := response.(model.WebResponseWithData); ok {
		data, _ := controller.UserService.FindUserByUsername(r.Context(), req.Username)
		helper.SetCokkie("role", strings.ToLower(data.Role), w)
		helper.SetCokkie("id", strconv.Itoa(data.ID), w)
		helper.WriteToResponseBody(w, response, nil, http.StatusOK)
		return
	}
	helper.WriteToResponseBody(w, response, err, http.StatusBadRequest)
}

func (controller *InitController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	req := model.UserUpdate{}
	id := helper.GetCokkie("id", r)
	helper.ReadFromRequestBody(r, &req)
	err := controller.UserService.UpdateUser(r.Context(), req, id)
	response := helper.ResponseWithMessage(err, "Success To Update User")
	helper.WriteToResponseBody(w, response, err, http.StatusBadRequest)
}
func (controller *InitController) FindAllUsers(w http.ResponseWriter, r *http.Request) {
	data, err := controller.UserService.FindAllUser(r.Context())
	response := helper.ResponseWithData(err, data)
	helper.WriteToResponseBody(w, response, err, http.StatusBadRequest)
}
func (controller *InitController) FindByUserid(w http.ResponseWriter, r *http.Request) {

	cusid := helper.GetCokkie("id", r)
	data, err := controller.UserService.FindUserById(r.Context(), cusid)
	response := helper.ResponseWithData(err, data)
	helper.WriteToResponseBody(w, response, err, http.StatusBadRequest)
}

func (controller *InitController) Logout(w http.ResponseWriter, r *http.Request) {

	go r.Header.Set("Authorization", "")
	go helper.SetCokkie("id", "", w)
	go helper.SetCokkie("role", "", w)
	response := make(chan model.WebResponseWithMessage)
	defer close(response)
	go func() {
		response <- helper.ResponseWithMessage(nil, "Berhasil Logout")
	}()
	helper.WriteToResponseBody(w, <-response, nil, http.StatusOK)
}
