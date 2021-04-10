package main

import (
	"github.com/intellij_test_golang/bd"
	"github.com/intellij_test_golang/handlers"
	"log"
)

func main ()  {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("sin conexion")
		return
	}


	handlers.Manejadores()
}



