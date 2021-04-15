package routers

import (
	"encoding/json"
	"github.com/intellij_test_golang/bd"
	"net/http"
)

// VerPerfil permite devolver los valores de un perfil determinado
func VerPerfil(writer http.ResponseWriter, request *http.Request) {

	ID := request.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(writer, "Se debe enviar el parametro ID", http.StatusBadRequest)
		return
	}
	perfil, err := bd.BuscoPerfil(ID)

	if err != nil {
		http.Error(writer, "Ocurrio un error al buscar el registro", http.StatusBadRequest)
		return
	}

	writer.Header().Set("context-type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(perfil)

}
