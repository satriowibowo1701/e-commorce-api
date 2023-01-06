package execption

import (
	"fmt"
	"net/http"

	"github.com/satriowibowo1701/e-commorce-api/helper"
	"github.com/satriowibowo1701/e-commorce-api/model"
)

func NotAllowed(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"Message":"Not Allowed"}`)
}

func UnAuthorized(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusUnauthorized)
	messages := &model.WebResponseWithMessage{Status: 401,
		Message: message}
	helper.WriteToResponseBody(w, messages)
}
