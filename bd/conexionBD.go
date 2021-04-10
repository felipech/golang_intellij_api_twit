package bd

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var MongoCN = ConectarBD()


var clientOptions = options.Client().ApplyURI("mongodb+srv://felipeMongoDb:FelipeMongoDb%21%22%23@cluster0.qxf07.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

// ConectarBD devuelve un objeto del tipo *mongo.Client
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("conexion exitosa con la BD")
	//si pasa las comprobaciones entonces hay una conexion valida
	return client
}

//ChequeoConnection es el ping a la bd
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}