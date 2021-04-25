package routers

import (
	"github.com/intellij_test_golang/bd"
	"net/http"
)

// EliminarTweet elimina un tweet de la bd
func EliminarTweet(writer http.ResponseWriter, request *http.Request) {

	ID := request.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(writer, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	err := bd.BorroTweet(ID, IDUsuario)
	if err != nil {
		http.Error(writer, "Ocurrio un error al borrar el tweet"+err.Error(), http.StatusBadRequest)
		return
	}
	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(http.StatusCreated)

}
