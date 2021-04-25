package routers

import (
	"github.com/intellij_test_golang/bd"
	"github.com/intellij_test_golang/models"
	"io"
	"net/http"
	"os"
	"strings"
)

// SubirAvatar metodo para subir una imagen
func SubirAvatar(write http.ResponseWriter, request *http.Request) {

	file, handler, err := request.FormFile("avatar")

	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/avatars/" + IDUsuario + "." + extension

	//0666 permisos de lectura y escritura
	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(write, "Error al subir la imagen "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(write, "Error al copiar la imagen "+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Avatar = IDUsuario + "." + extension
	status, err = bd.ModificoRegistro(usuario, IDUsuario)

	if err != nil || status == false {
		http.Error(write, "Error al grabar el avatar en la BD "+err.Error(), http.StatusBadRequest)
		return
	}

	write.Header().Set("Content-Type", "application/json")
	write.WriteHeader(http.StatusCreated)

}
