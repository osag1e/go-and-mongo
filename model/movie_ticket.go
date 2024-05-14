package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type MovieTicket struct {
	ID    primitive.ObjectID `bson:"id" json:"id"`
	Title string             `bson:"title" json:"title"`
	Price float64            `bson:"price" json:"price"`
}
