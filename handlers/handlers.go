package handlers

import (
	"github.com/gorilla/mux"
	"github.com/intellij_test_golang/middlew"
	"github.com/intellij_test_golang/routers"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)


//Manejadores seteo mi puerto
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
