package middleware

import (
	"net/http"

	"github.com/satriowibowo1701/e-commorce-api/execption"
	"github.com/satriowibowo1701/e-commorce-api/helper"
	"github.com/satriowibowo1701/e-commorce-api/router"
)

func AuthtenticationMiddleware(router *router.Method) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if router.Middleware[r.URL.Path] != "noauth" {
			err := helper.Authentication(r)

			if err != nil {
				execption.UnAuthorized(w, err.Error())
				return
			}
			data, err2 := helper.ClaimsAuthToken(r)
			if err2 != nil {
				execption.UnAuthorized(w, err2.Error())
				return
			}
			if err3 := helper.Authorization(data.Role, data.Id, r); err3 != nil {
				execption.UnAuthorized(w, err3.Error())
				return
			}

			if router.Middleware[r.URL.Path] != data.Role {
				execption.UnAuthorized(w, "Unauthorized")
				return
			}

		}
		router.ServeHTTP(w, r)

	})
}
