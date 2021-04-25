package routers

import (
	"github.com/intellij_test_golang/bd"
	"github.com/intellij_test_golang/models"
	"net/http"
)

func BajarRelacion(write http.ResponseWriter, request *http.Request) {
	ID := request.URL.Query().Get("id")

	var t models.Relacion

	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)
	if err != nil {
		http.Error(write, "Ocurrio un error al borrar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(write, " No se a logrado borrar la relacion", http.StatusBadRequest)
		return
	}

	write.WriteHeader(http.StatusCreated)

}
