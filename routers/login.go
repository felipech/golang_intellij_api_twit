package routers

import (
	"encoding/json"
	"github.com/intellij_test_golang/bd"
	"github.com/intellij_test_golang/jwt"
	"github.com/intellij_test_golang/models"
	"net/http"
	"time"
)

func Login(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Add("content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(writer, "Usuario y/o contraceña invalidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(writer, "El email es necesario", 400)
		return
	}

	documento, existe := bd.IntentoLogin(t.Email, t.Password)

	if existe == false {
		http.Error(writer, "Usuario y/o contraseña invalidos", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)

	if err != nil {
		http.Error(writer, "Ocurrio un error al generar el Token"+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)

	http.SetCookie(writer, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
