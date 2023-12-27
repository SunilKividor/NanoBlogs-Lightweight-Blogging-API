package models

import (
	"time"

	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	First_name string             `json:"first_name" validate:"required,min=3,max=30"`
	Last_name  string             `json:"last_name" validate:"required,min=3,max=30"`
	Password   string             `json:"password" validate:"required,min=6"`
	Email      string             `json:"email" validate:"email,required"`
	Tags       []string           `json:"tags" validate:"required"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}

func GetUser() *User {
	return &User{}
}

func (u *User) ValidateUser() error {
	validate := validator.New()
	return validate.Struct(u)
}
