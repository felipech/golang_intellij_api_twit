package routers

import (
	"encoding/json"
	"github.com/intellij_test_golang/bd"
	"net/http"
	"strconv"
)

func LeoTweetsSeguidores(write http.ResponseWriter, request *http.Request) {

	if len(request.URL.Query().Get("pagina")) < 1 {
		http.Error(write, "debe enviar el parametro de la pagina", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(request.URL.Query().Get("pagina"))

	if err != nil {
		http.Error(write, "Debe enviar el parametro como entero mayor a 0", http.StatusBadRequest)
		return
	}
	respuesta, correcto := bd.LeoTweetSeguidores(IDUsuario, pagina)
	if correcto == false {
		http.Error(write, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	write.Header().Set("Content-Type", "application/json")
	json.NewEncoder(write).Encode(respuesta)

}
