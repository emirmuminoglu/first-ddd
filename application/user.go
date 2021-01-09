package application

import (
	"context"

	"github.com/emirmuminoglu/first-ddd/domain/repository"
	"github.com/emirmuminoglu/first-ddd/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserInteractor struct {
	Repository repository.UserRepository
}

func (i UserInteractor) GetUser(ctx context.Context, id primitive.ObjectID) (*domain.User, error) {
	return i.Repository.Get(ctx, id)
}

func (i UserInteractor) CreateUser(ctx context.Context, name string) error {
	u, err := domain.NewUser(name)
	if err != nil {
		return err
	}

	return i.Repository.Save(ctx, u)
}