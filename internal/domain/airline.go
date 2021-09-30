package domain

import "go.mongodb.org/mongo-driver/mongo"

type Airline interface {
	Feed(setting *mongo.SingleResult) []byte
}
