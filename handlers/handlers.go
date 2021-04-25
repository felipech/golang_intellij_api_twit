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
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoBD(routers.GraboTweet)).Methods("POST")
	router.HandleFunc("/leoTweets", middlew.ChequeoBD(routers.Leotweets)).Methods("GET")
	router.HandleFunc("/borroTweet", middlew.ChequeoBD(routers.EliminarTweet)).Methods("DELETE")
	router.HandleFunc("/subirAvatar", middlew.ChequeoBD(routers.SubirAvatar)).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlew.ChequeoBD(routers.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/subirBanner", middlew.ChequeoBD(routers.SubirBanner)).Methods("POST")
	router.HandleFunc("/obtenerBanner", middlew.ChequeoBD(routers.ObtenerBanner)).Methods("GET")
	router.HandleFunc("/altaRelacion", middlew.ChequeoBD(routers.Altarelacion)).Methods("POST")
	router.HandleFunc("/bajaRelacion", middlew.ChequeoBD(routers.BajarRelacion)).Methods("DELETE")
	router.HandleFunc("/consultaRelacion", middlew.ChequeoBD(routers.ConsultaRelacion)).Methods("GET")
	router.HandleFunc("/listaUsuarios", middlew.ChequeoBD(routers.ListaUsuarios)).Methods("GET")
	router.HandleFunc("/leoTweetSeguidores", middlew.ChequeoBD(routers.LeoTweetsSeguidores)).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
