package domain

import "go.mongodb.org/mongo-driver/mongo"

type AirlineSetting interface {
	FindAllProviders() map[string]*mongo.Cursor
}
