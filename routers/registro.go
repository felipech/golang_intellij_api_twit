package routers

import (
	"encoding/json"
	"github.com/intellij_test_golang/bd"
	"github.com/intellij_test_golang/models"
	"net/http"
)

func Registro(writer http.ResponseWriter, request * http.Request){
	var t models.Usuario

	err := json.NewDecoder(request.Body).Decode(&t)

	if err != nil {
		http.Error(writer, "Errpr en los datos", 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(writer, "El correo es un campo obligatorio", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(writer, "la password debe tener minimo de 6 caracteres", 400)
		return
	}
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)

	if encontrado {
		http.Error(writer, "Ya existe el usuario en la bd", 400)
		return
	}
	_,status, err := bd.InsertoRegistro(t)

	if err != nil{
		http.Error(writer, "Ocurrio un error al insertar el usuario" + err.Error(),400)
		return
	}

	if status == false{
		http.Error(writer, "No se logro insertar el registro del usuario",400)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}
