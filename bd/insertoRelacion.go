package bd

import (
	"context"
	"github.com/intellij_test_golang/models"
	"time"
)

// InsertoRelacion inserta la relacion en la bd
func InsertoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twitter_golang")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, t)

	if err != nil {
		return false, err
	}

	return true, nil

}
