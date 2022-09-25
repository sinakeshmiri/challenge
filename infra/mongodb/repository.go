package mongodb

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	client     *mongo.Client
	db         string
	collection string
}

func newStorage() Storage {
	return Storage{
		client:     NewClient(),
		db:         os.Getenv("DATABASE"),
		collection: os.Getenv("COLLECTION"),
	}
}

func NewClient() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		log.Fatalf("mongodb new client error: %v", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("mongodb connect error: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("mongodb ping error: %v", err)
	}

	return client
}

type Repository interface {
	insert(entity interface{}) error
}

func (s Storage) insert(entity interface{}) error {
	c := newStorage()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := c.client.Database(c.db).Collection(c.collection).InsertOne(ctx, entity)

	return err
}
