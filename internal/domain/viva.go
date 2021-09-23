package domain

type VivaSetting struct {
	Origin       string         `json:"origin"`
	BeginDate    string         `json:"beginDate"`
	Passengers   PassengersViva `json:"passengers"`
	Destination  string         `json:"destination"`
	CurrencyCode string         `json:"currencyCode"`
	EndDate      string         `json:"endDate"`
	PromoCode    string         `json:"promoCode"`
}

type PassengersViva struct {
	Adt int `json:"ADT"`
}
