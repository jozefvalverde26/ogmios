package latam

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/jozefvalverde26/ogmios/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	macro = "?sort=%s&cabinType=%s&origin=%s&destination=%s&inFlightDate=%s&inFrom=%s&inOfferId=%s&outFlightDate=%s&outFrom=%s&outOfferId=%s&adult=%s&child=%s&infant=%s&redemption=%s"
)

type Service struct {
	config Config
}

func NewService(config Config) Service {
	return Service{config}
}

func (s Service) Feed(cur *mongo.SingleResult) []byte {
	var setting domain.LatamSetting
	err := cur.Decode(&setting)
	if err != nil {
		log.Fatal(err)
	}
	URL := fmt.Sprintf(
		fmt.Sprintf("%s%s", s.config.FeedURL, macro),
		setting.Sort,
		setting.CabinType,
		setting.Origin,
		setting.Destination,
		setting.InFlightDate,
		setting.InFrom,
		setting.InOfferID,
		setting.OutFlightDate,
		setting.OutFrom,
		setting.OutOfferId,
		setting.Adult,
		setting.Child,
		setting.Infant,
		setting.Redemption,
	)
	reqURL, _ := url.Parse(URL)
	headers := map[string][]string{
		"Content-Type":                {"application/json; charset=UTF-8"},
		"x-latam-app-session-id":      {"959ff346-3dc4-4f0a-ad0d-2f7b0bcaf6d1"},
		"x-latam-application-country": {"PE"},
		"x-latam-application-lang":    {"es"},
		"x-latam-application-name":    {"web-air-offers"},
		"x-latam-application-oc":      {"pe"},
		"x-latam-client-name":         {"web-air-offers"},
		"x-latam-request-id":          {"e1ceaac2-4318-4073-ad71-3776bfcee2a5"},
		"x-latam-track-id":            {"6e36b10a-4125-4725-b563-e6dc3881ec1b"},
	}
	req := &http.Request{
		Method: "GET",
		URL:    reqURL,
		Header: headers,
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	//read response data
	data, _ := ioutil.ReadAll(res.Body)

	//close response body
	res.Body.Close()

	return data
}
