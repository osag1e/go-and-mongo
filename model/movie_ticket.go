package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type MovieTicket struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title string             `bson:"title" json:"title"`
	Price float64            `bson:"price" json:"price"`
}
