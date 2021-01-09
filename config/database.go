package config

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDBConnection() (*mongo.Client, error) {
	uri := os.Getenv("DB_URI")

	opts := options.Client().ApplyURI(uri)

	return mongo.Connect(context.Background(), opts)
}
