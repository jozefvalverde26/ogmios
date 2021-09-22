package domain

type AirlineSetting interface {
	FindAllProviders() map[string]SkySetting
}
