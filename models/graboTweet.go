package models

import "time"

// GraboTweet estructura para grabar el tweet
type GraboTweet struct {
	UserID  string    `bson:"user_id" json:"user_id,omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
