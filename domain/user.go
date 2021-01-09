package domain

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrInvalidName = errors.New("invalid name")
)

//easyjson:json
type User struct {
	ID   primitive.ObjectID `json:"_id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
}

func NewUser(name string) (*User, error) {
	if name == "" {
		return nil, ErrInvalidName
	}

	return &User{
		ID:   primitive.NewObjectID(),
		Name: name,
	}, nil
}
