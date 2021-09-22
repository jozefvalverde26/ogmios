package mongo

import (
	"context"
	"log"
	"time"

	"github.com/jozefvalverde26/ogmios/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service struct {
	collection *mongo.Collection
}

func NewService(config Config) Service {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Uri))
	if err != nil {
		panic(err)
	}
	collection := client.Database(config.Db).Collection(config.Collection)

	return Service{collection}
}

func (s Service) FindAllProviders() map[string]domain.SkySetting {
	mappper := make(map[string]domain.SkySetting)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, curError := s.collection.Find(ctx, bson.D{})
	if curError != nil {
		panic(curError)
	}
	for cur.Next(ctx) {
		var setting domain.SkySetting
		err := cur.Decode(&setting)
		if err != nil {
			log.Fatal(err)
		}
		mappper[setting.Provider] = setting
	}
	return mappper
}
