package collector

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/jozefvalverde26/ogmios/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
)

type Service struct {
	setting   domain.AirlineSetting
	providers map[string]domain.Airline
}

func NewService(setting domain.AirlineSetting, providers map[string]domain.Airline) Service {
	return Service{setting, providers}
}

func (s Service) Process() {
	fmt.Println("Collect process started...")
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(len(s.providers))
	for name, provider := range s.providers {
		go s.collect(name, provider, &wg)
	}
	wg.Wait()
	fmt.Printf("%s %v\n", "Collect process finished after: ", time.Since(start))
}

func (s Service) collect(name string, provider domain.Airline, wg *sync.WaitGroup) {
	defer wg.Done()
	airlineSetting := s.setting.FindProviderByName(name)
	var bsonData bson.D
	data := provider.Feed(airlineSetting)
	bson.UnmarshalExtJSON(data, true, &bsonData)
	err := s.setting.Saveflight(bsonData)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%s: %v\n", name, err)
}
