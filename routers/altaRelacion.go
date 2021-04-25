package routers

import (
	"github.com/intellij_test_golang/bd"
	"github.com/intellij_test_golang/models"
	"net/http"
)

// Altarelacion inserta la relacion entre usuarios
func Altarelacion(write http.ResponseWriter, request *http.Request) {

	ID := request.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(write, "El ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)
	if err != nil {
		http.Error(write, "Ocurrio un error al insertar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(write, " No se a logrado insertar la relacion", http.StatusBadRequest)
		return
	}

	write.WriteHeader(http.StatusCreated)
}
