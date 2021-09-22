package domain

type SkySetting struct {
	Provider           string     `json:"provider"`
	Segments           []Segments `json:"segments"`
	Passengers         Passengers `json:"passengers"`
	Currency           string     `json:"currency"`
	DepartureDateStart string     `json:"departureDateStart"`
	DepartureDateEnd   string     `json:"departureDateEnd"`
}

type Segments struct {
	From    string `json:"from"`
	To      string `json:"to"`
	DepDate string `json:"depDate"`
}

type Passengers struct {
	Adults   int `json:"adults"`
	Children int `json:"children"`
	Babies   int `json:"babies"`
}
