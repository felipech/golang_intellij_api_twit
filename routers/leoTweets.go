package routers

import (
	"encoding/json"
	"github.com/intellij_test_golang/bd"
	"net/http"
	"strconv"
)

func Leotweets(writer http.ResponseWriter, request *http.Request) {

	ID := request.URL.Query().Get("id")

	if len(ID) < 0 {
		http.Error(writer, "Debe enviar el parametro", http.StatusBadRequest)
		return
	}

	if len(request.URL.Query().Get("pagina")) < 1 {
		http.Error(writer, "Debe enviar el parametro pagina ", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(request.URL.Query().Get("pagina"))

	if err != nil {
		http.Error(writer, "Debe enviar el parametro con un valor mayor a cero", http.StatusBadRequest)
		return
	}

	pag := int64(pagina)

	respuesta, correcto := bd.LeoTweets(ID, pag)

	if correcto == false {
		http.Error(writer, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(respuesta)

}
