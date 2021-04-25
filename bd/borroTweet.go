package bd

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// BorroTweet borra un tweet de la bd
func BorroTweet(ID string, UserID string) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twitter_golang")
	col := db.Collection("tweet")

	obgID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id":    obgID,
		"userid": UserID,
	}
	_, err := col.DeleteOne(ctx, condicion)
	return err
}
