package bd

import (
	"context"
	"github.com/intellij_test_golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// LeoTweets funcion para devolver la lista de tweet de un usuario
func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("twitter_golang")
	col := db.Collection("tweet")

	//slice de tweet de cada usuario
	var resultados []*models.DevuelvoTweets

	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find()
	opciones.SetLimit(20)
	//-1 es orden desendente
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	//este es un calculo que va a decidir cuantos paginas se debe saltar cada vez que se hace la solicitud
	//si es la pimera pagina se salta cero y si va aumentando entonces va saltando las 20 primeras despues las 40 y asi
	opciones.SetSkip((pagina - 1) * 20)

	cursor, err := col.Find(ctx, condicion, opciones)

	//si no encuentra nada entonces devolvemos el error
	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}

	for cursor.Next(context.TODO()) {
		var registro models.DevuelvoTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}

		resultados = append(resultados, &registro)
	}

	return resultados, true

}
