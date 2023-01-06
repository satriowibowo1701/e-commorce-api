package helper

import (
	"net/http"
	"strconv"

	"github.com/satriowibowo1701/e-commorce-api/model"
)

func GetParam(key string, r *http.Request) int {

	res, err := strconv.Atoi(r.URL.Query().Get(key))
	if err != nil {
		return -1

	}
	return res

}

func ResponseWithMessage(err error, psn string) model.WebResponseWithMessage {

	var message string
	if err != nil {
		message = err.Error()
		response := model.WebResponseWithMessage{
			Status:  http.StatusBadRequest,
			Message: message,
		}
		return response

	}
	response := model.WebResponseWithMessage{
		Status:  http.StatusOK,
		Message: psn,
	}
	return response

}

func ResponseWithData(err error, data interface{}) interface{} {

	if err != nil {

		response := model.WebResponseWithMessage{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		return response

	}

	response := model.WebResponseWithData{
		Status: http.StatusOK,
		Data:   data,
	}
	return response

}

func SetCokkie(key, value string, w http.ResponseWriter) {
	cookie := &http.Cookie{}
	cookie.Name = key
	cookie.Value = value
	cookie.Path = "/"
	http.SetCookie(w, cookie)
}

func GetCokkie(key string, r *http.Request) int {
	cookie, err := r.Cookie(key)
	if err != nil {
		return -1
	}
	newcookie, _ := strconv.Atoi(cookie.Value)
	return newcookie
}

// func GenereteCSRF(Id int, r *http.Request) string {

// }
