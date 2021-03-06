package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/jozefvalverde26/ogmios/internal/collector"
	"github.com/jozefvalverde26/ogmios/internal/domain"
	"github.com/jozefvalverde26/ogmios/internal/latam"
	"github.com/jozefvalverde26/ogmios/internal/mongo"
	"github.com/jozefvalverde26/ogmios/internal/sky"
	"github.com/jozefvalverde26/ogmios/internal/viva"
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

	vivaConfig := viva.Config{
		FeedURL: os.Getenv("VIVA_FEED_URL"),
	}
	vivaService := viva.NewService(vivaConfig)

	latamConfig := latam.Config{
		FeedURL: os.Getenv("LATAM_FEED_URL"),
	}
	latamService := latam.NewService(latamConfig)
	mongoConfig := mongo.Config{
		Uri: os.Getenv("MONGO_URI"),
		Db:  os.Getenv("MONGO_DB_NAME"),
	}
	mongoService := mongo.NewService(mongoConfig)

	providers := map[string]domain.Airline{
		"sky":   skyService,
		"viva":  vivaService,
		"latam": latamService,
	}
	collectorService := collector.NewService(mongoService, providers)
	collectorService.Process()
}
