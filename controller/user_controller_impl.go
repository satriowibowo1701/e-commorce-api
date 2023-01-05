package controller

import (
	"net/http"
	"strconv"

	"github.com/satriowibowo1701/e-commorce-api/helper"
	"github.com/satriowibowo1701/e-commorce-api/model"
)

func (controller *InitController) RegisterUser(w http.ResponseWriter, r *http.Request) {

	req := model.UserRegis{}
	helper.ReadFromRequestBody(r, &req)
	err := controller.UserService.CreateUser(r.Context(), req)
	response := helper.ResponseWithMessage(err, "Succes Creating User")
	helper.WriteToResponseBody(w, response)
}

func (controller *InitController) LoginUser(w http.ResponseWriter, r *http.Request) {

	req := model.LoginRequest{}
	helper.ReadFromRequestBody(r, &req)
	token, err := controller.UserService.Login(r.Context(), req)
	response := helper.ResponseWithData(err, token)
	if _, ok := response.(model.WebResponseWithData); ok {
		data, _ := controller.UserService.FindUserByUsername(r.Context(), req.Username)
		helper.SetCokkie("id", strconv.Itoa(data.ID), w)
		helper.SetCokkie("role", data.Role, w)
		helper.WriteToResponseBody(w, response)
		return
	}
	helper.WriteToResponseBody(w, response)
}

func (controller *InitController) UpdateUser(w http.ResponseWriter, r *http.Request) {

	req := model.UserUpdate{}
	go helper.ReadFromRequestBody(r, &req)
	err := controller.UserService.UpdateUser(r.Context(), req)
	response := helper.ResponseWithMessage(err, "Success To Update User")
	helper.WriteToResponseBody(w, response)
}
func (controller *InitController) FindAllUsers(w http.ResponseWriter, r *http.Request) {
	data, err := controller.UserService.FindAllUser(r.Context())
	response := helper.ResponseWithData(err, data)
	helper.WriteToResponseBody(w, response)
}
func (controller *InitController) FindByUserid(w http.ResponseWriter, r *http.Request) {
	req := model.User{}
	helper.ReadFromRequestBody(r, &req)
	data, err := controller.UserService.FindUserById(r.Context(), int64(req.ID))
	response := helper.ResponseWithData(err, data)
	helper.WriteToResponseBody(w, response)
}
