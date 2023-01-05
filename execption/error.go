package execption

import (
	"fmt"
	"net/http"
)

func NotAllowed(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"Message":"Not Allowed"}`)
}
