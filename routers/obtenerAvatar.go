package routers

import (
	"github.com/intellij_test_golang/bd"
	"io"
	"net/http"
	"os"
)

// ObtenerAvatar otiene el avatar del un usuario
func ObtenerAvatar(write http.ResponseWriter, request *http.Request) {

	ID := request.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(write, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)

	if err != nil {
		http.Error(write, "Usuario no encontrado", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/avatars" + perfil.Avatar)
	if err != nil {
		http.Error(write, "Imagen no encontrado", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(write, OpenFile)

	if err != nil {
		http.Error(write, "Error al copiar la imagen", http.StatusBadRequest)
	}

}
