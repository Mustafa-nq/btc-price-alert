package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Alert struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Coin      string             `json:"coin" bson:"coin"`
	Price     float64            `json:"price" bson:"price"`
	Direction string             `json:"direction" bson:"direction"`
	Email     string             `json:"email" bson:"email"`
}
