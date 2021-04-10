package bd

import (
	"context"
	"github.com/intellij_test_golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// ChequeoYaExisteUsuario chquea si existe un usuario en la bd
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string){

	ctx, cancel := context.WithTimeout(context.Background(), 15* time.Second)
	defer cancel()

	db := MongoCN.Database("twitter_golang")
	col := db.Collection("usuarios")

	//es formato JSON por eso usa {}
	condicion := bson.M{"email":email}

	var resultado models.Usuario
	err := col.FindOne(ctx, condicion).Decode(&resultado)

	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID

}