package application

import (
	"context"

	"errors"
	"github.com/emirmuminoglu/first-ddd/domain"
	"github.com/emirmuminoglu/first-ddd/domain/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrDuplicateKey = errors.New("duplicate key error")
)

type UserInteractor struct {
	Repository repository.UserRepository
}

func (i UserInteractor) GetUser(ctx context.Context, email string) (*domain.User, error) {
	return i.Repository.Get(ctx, email)
}

func isDup(err error) bool {
	var e mongo.WriteException
	if errors.As(err, &e) {
		for _, we := range e.WriteErrors {
			if we.Code == 11000 {
				return true
			}
		}
	}
	return false
}

func (i UserInteractor) CreateUser(ctx context.Context, email, password string) error {
	u, err := domain.NewUser(email, password)
	if err != nil {
		return err
	}

	err = i.Repository.Save(ctx, u)
	if err != nil {
		if isDup(err) {
			return ErrDuplicateKey
		}

		return err
	}
	return err
}
