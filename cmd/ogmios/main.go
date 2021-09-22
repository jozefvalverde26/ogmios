package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/jozefvalverde26/ogmios/internal/collector"
	"github.com/jozefvalverde26/ogmios/internal/domain"
	"github.com/jozefvalverde26/ogmios/internal/mongo"
	"github.com/jozefvalverde26/ogmios/internal/sky"
)

func main() {
	// load .env file from given path
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	skyConfig := sky.Config{
		FeedURL:   os.Getenv("SKY_FEED_URL"),
		FeedToken: os.Getenv("SKY_FEED_TOKEN"),
	}
	skyService := sky.NewService(skyConfig)

	mongoConfig := mongo.Config{
		Uri:        os.Getenv("MONGO_URI"),
		Db:         os.Getenv("MONGO_DB_NAME"),
		Collection: os.Getenv("MONGO_COLLECTION_NAME"),
	}
	mongoService := mongo.NewService(mongoConfig)

	providers := []domain.Airline{skyService}
	collectorService := collector.NewService(mongoService, providers)
	collectorService.Process()
}
