package mongo

import (
	"context"
	"log"
	"time"

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

type Result struct {
	Provider string `json:"provider"`
}

func (s Service) FindAllProviders() map[string]*mongo.Cursor {
	mappper := make(map[string]*mongo.Cursor)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, curError := s.collection.Find(ctx, bson.D{})
	if curError != nil {
		panic(curError)
	}
	for cur.Next(ctx) {
		var res Result
		err := cur.Decode(&res)
		if err != nil {
			log.Fatal(err)
		}
		mappper[res.Provider] = cur
	}
	return mappper
}
