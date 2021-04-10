package middlew

import (
	"github.com/intellij_test_golang/bd"
	"net/http"
)

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc{

	return func(writer http.ResponseWriter, request *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(writer,"Conexion perdida con la base de datos", 500 )
			return
		}
		next.ServeHTTP(writer, request)
	}
}
