package domain

type LatamSetting struct {
	Sort          string `json:"sort"`
	CabinType     string `json:"cabinType"`
	Origin        string `json:"origin"`
	Destination   string `json:"destination"`
	InFlightDate  string `json:"inFlightDate"`
	InFrom        string `json:"inFrom"`
	InOfferID     string `json:"inOfferId"`
	OutFlightDate string `json:"outFlightDate"`
	OutFrom       string `json:"outFrom"`
	OutOfferId    string `json:"outOfferId"`
	Adult         string `json:"adult"`
	Child         string `json:"child"`
	Infant        string `json:"infant"`
	Redemption    string `json:"redemption"`
}
