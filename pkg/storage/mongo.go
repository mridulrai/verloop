package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Storage implements payment.Repository interface.
type Storage struct {
	Client *mongo.Client
	DB     *mongo.Database
}

// NewStorage is a factory function to create a new Storage.
func NewStorage(mongoURI string, dbname string) (*Storage, error) {
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	db := client.Database(dbname)
	s := &Storage{Client: client, DB: db}
	return s, nil
}
