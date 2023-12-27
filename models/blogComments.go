package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type BlogComments struct {
	ID       primitive.ObjectID `json:"_id,omitempty"`
	BlogId   primitive.ObjectID `json:"blog_id" validate:"required"`
	Comments []Comment          `json:"comments"`
}
