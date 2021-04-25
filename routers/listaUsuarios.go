package routers

import (
	"encoding/json"
	"github.com/intellij_test_golang/bd"
	"net/http"
	"strconv"
)

func ListaUsuarios(write http.ResponseWriter, request *http.Request) {

	typeUser := request.URL.Query().Get("type")
	page := request.URL.Query().Get("page")
	search := request.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)

	if err != nil {
		http.Error(write, "Deve enviar el parametro pagina como entero mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp)

	result, status := bd.LeoUsuariosTodos(IDUsuario, pag, search, typeUser)

	if status == false {
		http.Error(write, "Error al leer los usuarios", http.StatusBadRequest)
		return
	}

	write.Header().Set("Content-Type", "application/json")
	write.WriteHeader(http.StatusCreated)
	json.NewEncoder(write).Encode(result)

}
