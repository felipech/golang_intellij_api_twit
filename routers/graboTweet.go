package routers

import (
	"encoding/json"
	"github.com/intellij_test_golang/bd"
	"github.com/intellij_test_golang/models"
	"net/http"
	"time"
)

func GraboTweet(write http.ResponseWriter, request *http.Request) {

	var mensaje models.GraboTweet

	err := json.NewDecoder(request.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  mensaje.UserID,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)

	if err != nil {
		http.Error(write, "Error al grabar el tweet "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(write, "No se ha logrado insertar el tweet ", 400)
		return
	}

	write.WriteHeader(http.StatusCreated)

}
