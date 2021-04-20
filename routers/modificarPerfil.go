package routers

import (
	"encoding/json"
	"github.com/intellij_test_golang/bd"
	"github.com/intellij_test_golang/models"
	"net/http"
)

func ModificarPerfil(writer http.ResponseWriter, request *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(request.Body).Decode(&t)

	if err != nil {
		http.Error(writer, "Datos Incorrectos "+err.Error(), 400)
		return
	}

	var status bool
	status, err = bd.ModificoRegistro(t, IDUsuario)

	if err != nil {
		http.Error(writer, "Ocurrio un error al intetar modificar el registro. Reintentar nuevamente"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(writer, "No se pudo modificar el registro de usuario ", 400)
		return
	}

	writer.WriteHeader(http.StatusCreated)

}
