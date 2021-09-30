package domain

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AirlineSetting interface {
	Saveflight(flightBson bson.D) error
	FindProviderByName(name string) *mongo.SingleResult
}
