package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	UserId      primitive.ObjectID `json:"user_id" bson:"user_id" validate:"required" `
	UserComment string             `json:"user_comment" bson:"user_comment" validate:"required"`
	CreatedAt   time.Time          `json:"created_at"`
}
