package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}, err error, httpcode int) {
	writer.Header().Add("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(httpcode)
	}
	encoder := json.NewEncoder(writer)
	_ = encoder.Encode(response)

}
func WriteToResponseLogin(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}
