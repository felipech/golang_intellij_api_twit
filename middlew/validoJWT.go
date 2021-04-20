package middlew

import (
	"github.com/intellij_test_golang/routers"
	"net/http"
)

func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {

	return func(writer http.ResponseWriter, request *http.Request) {
		_, _, _, err := routers.ProcesoToken(request.Header.Get("Authorization"))

		if err != nil {
			http.Error(writer, "error en el token"+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(writer, request)
	}
}
