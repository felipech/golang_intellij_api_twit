package bd

import (
	"context"
	"fmt"
	"github.com/intellij_test_golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func LeoUsuariosTodos(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twitter_golang")
	col := db.Collection("tweet")

	var results []*models.Usuario

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, findOptions)

	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	var encontrado, incluir bool

	for cur.Next(ctx) {
		var su models.Usuario
		err := cur.Decode(&su)

		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}
		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = su.ID.Hex()

		incluir = false

		encontrado, err = ConsultoRelacion(r)

		if tipo == "new" && encontrado == false {
			incluir = true
		}

		if tipo == "follow" && encontrado == true {
			incluir = true
		}

		if r.UsuarioRelacionID == ID {
			incluir = false
		}

		if incluir == true {
			su.Password = ""
			su.Biografia = ""
			su.SitioWeb = ""
			su.Ubicacion = ""
			su.Banner = ""
			su.Email = ""
			results = append(results, &su)
		}

	}

	err = cur.Err()

	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	cur.Close(ctx)
	return results, true
}
