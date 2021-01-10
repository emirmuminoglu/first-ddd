package repository

import (
	"context"

	"github.com/emirmuminoglu/first-ddd/domain"
)

type UserRepository interface {
	Get(ctx context.Context, email string) (*domain.User, error)
	Save(ctx context.Context, user *domain.User) error
}
