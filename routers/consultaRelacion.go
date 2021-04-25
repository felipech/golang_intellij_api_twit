package routers

import (
	"encoding/json"
	"github.com/intellij_test_golang/bd"
	"github.com/intellij_test_golang/models"
	"net/http"
)

func ConsultaRelacion(write http.ResponseWriter, request *http.Request) {

	ID := request.URL.Query().Get("id")

	var t models.Relacion

	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	var resp models.RespuestaConsultaRelacion

	status, err := bd.ConsultoRelacion(t)

	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	write.Header().Set("Content-Type", "application/json")
	write.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(write).Encode(resp)
	if err != nil {
		return
	}

}
