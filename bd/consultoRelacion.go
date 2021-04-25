package bd

import (
	"context"
	"fmt"
	"github.com/intellij_test_golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// ConsultoRelacion consulta la relacion entre 2 usuarios
func ConsultoRelacion(t models.Relacion) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twitter_golang")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioID,
	}

	var resultado models.Relacion

	fmt.Println(resultado)

	err := col.FindOne(ctx, condicion).Decode(&resultado)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil

}
