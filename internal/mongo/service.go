package mongo

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service struct {
	db *mongo.Database
}

func NewService(config Config) Service {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Uri))
	if err != nil {
		panic(err)
	}
	db := client.Database(config.Db)

	return Service{db}
}

type Result struct {
	Provider string `json:"provider"`
}

func (s Service) FindProviderByName(name string) *mongo.SingleResult {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	collection := s.db.Collection(os.Getenv("MONGO_AIRLINE_SETTING_COLLECTION_NAME"))
	singleRes := collection.FindOne(ctx, bson.M{"provider": name})
	return singleRes
}

func (s Service) Saveflight(flightBson bson.D) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := s.db.Collection(os.Getenv("MONGO_FLIGHT_COLLECTION_NAME"))
	_, err := collection.InsertOne(ctx, flightBson)
	if err != nil {
		return err
	}
	return nil
}
