package sky

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/jozefvalverde26/ogmios/internal/domain"
)

type Service struct {
	config Config
}

func NewService(config Config) Service {
	return Service{config}
}

func (s Service) Feed(setting domain.SkySetting) {
	reqURL, _ := url.Parse(s.config.FeedURL)
	headers := map[string][]string{
		"Content-Type": {"application/json; charset=UTF-8"},
		"_token":       {s.config.FeedToken},
	}
	// body := ioutil.NopCloser(strings.NewReader(`
	// 	{
	// 		"segments": [
	// 			{
	// 				"from": "LIM",
	// 				"to": "AQP",
	// 				"depDate": "2021-09-21"
	// 			}
	// 		],
	// 		"passengers": {
	// 			"adults": 1,
	// 			"children": 0,
	// 			"babies": 0
	// 		},
	// 		"currency": "USD",
	// 		"departureDateStart": "2021-09-25",
	// 		"departureDateEnd": "2021-10-02"
	// 	}
	// `))
	jsonStr, _ := json.Marshal(setting)
	fmt.Printf("%s\n", jsonStr)
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

	fmt.Printf("%s\n", data)
}
