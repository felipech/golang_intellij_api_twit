package models

// Tweet body del tweet
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
