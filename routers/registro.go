package routers

import (
	"encoding/json"
	"github.com/intellij_test_golang/bd"
	"github.com/intellij_test_golang/models"
	"log"
	"net/http"
)

func Registro(writer http.ResponseWriter, request *http.Request) {
	var t models.Usuario

	//log.Println(json.NewDecoder(request.Body).Decode(&t))
	//cuando se usa el decode de un request body se tiene que cerrar la conexion sino va a dar "eof" al intentar decodificar la request de nuevo
	//cuidado con imprimir el resultado de esto
	////log.Println(json.NewDecoder(request.Body).Decode(&t))
	err := json.NewDecoder(request.Body).Decode(&t)
	log.Println("Nombre enviado en el JSON " + t.Nombre)

	if err != nil {
		http.Error(writer, "Error en los datos", 400)
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
	_, status, err := bd.InsertoRegistro(t)

	if err != nil {
		http.Error(writer, "Ocurrio un error al insertar el usuario"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(writer, "No se logro insertar el registro del usuario", 400)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}
