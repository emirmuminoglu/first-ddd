package repository

import (
	"context"

	"github.com/emirmuminoglu/first-ddd/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	Get(ctx context.Context, id primitive.ObjectID) (*domain.User, error)
	Save(ctx context.Context, user *domain.User) error
}
