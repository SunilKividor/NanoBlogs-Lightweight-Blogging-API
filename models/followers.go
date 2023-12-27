package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Follow struct {
	UserId    primitive.ObjectID   `json:"user_id"  bson:"user_id"`
	Followers []primitive.ObjectID `json:"followers" bson:"followers"`
	Following []primitive.ObjectID `json:"following" bson:"following"`
}
