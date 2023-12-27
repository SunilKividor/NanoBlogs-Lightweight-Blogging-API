package models

import (
	"time"

	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserId    primitive.ObjectID `json:"user_id" bson:"user_id" validate:"required"`
	Title     string             `json:"title"  bson:"title" validate:"required,min=3"`
	Body      string             `json:"body"  bson:"body" validate:"required,min=3"`
	Tags      []string           `json:"tags"  bson:"tags" validate:"required"`
	Timestamp time.Time          `json:"time"  bson:"time"`
}

func GetBlog() *Blog {
	return &Blog{
		Timestamp: time.Now(),
	}
}

func (b *Blog) ValidateBlog() error {
	validate := validator.New()
	return validate.Struct(b)
}
