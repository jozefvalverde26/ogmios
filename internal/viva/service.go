package viva

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/jozefvalverde26/ogmios/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	config Config
}

func NewService(config Config) Service {
	return Service{config}
}

func (s Service) Feed(cur *mongo.SingleResult) []byte {
	var setting domain.VivaSetting
	err := cur.Decode(&setting)
	if err != nil {
		log.Fatal(err)
	}
	reqURL, _ := url.Parse(s.config.FeedURL)
	headers := map[string][]string{
		"Content-Type": {"application/json; charset=UTF-8"},
	}
	jsonStr, _ := json.Marshal(setting)
	body := ioutil.NopCloser(strings.NewReader(string(jsonStr)))
	req := &http.Request{
		Method: "POST",
		URL:    reqURL,
		Header: headers,
		Body:   body,
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
