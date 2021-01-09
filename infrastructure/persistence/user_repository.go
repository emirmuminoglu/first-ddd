package persistence

import (
	"context"

	"github.com/emirmuminoglu/first-ddd/domain"
	"github.com/emirmuminoglu/first-ddd/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	dbname string
	client *mongo.Client
}

func NewUserRepository(client *mongo.Client, dbname string) repository.UserRepository {
	return &userRepository{client: client, dbname: dbname}
}

func (r *userRepository) Get(ctx context.Context, id primitive.ObjectID) (*domain.User, error) {
	col := r.col("user")

	u := new(domain.User)
	err := col.FindOne(ctx, bson.M{"_id": id}).Decode(u)

	return u, err
}

func (r *userRepository) Save(ctx context.Context, u *domain.User) error {
	col := r.col("user")

	_, err := col.InsertOne(ctx, u)

	return err
}

func (r *userRepository) col(colName string) *mongo.Collection {
	return r.client.Database(r.dbname).Collection(colName)
}
