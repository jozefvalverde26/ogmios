package collector

import (
	"github.com/jozefvalverde26/ogmios/internal/domain"
)

type Service struct {
	setting   domain.AirlineSetting
	providers []domain.Airline
}

func NewService(setting domain.AirlineSetting, providers []domain.Airline) Service {
	return Service{setting, providers}
}

func (s Service) Process() {
	airlineSettings := s.setting.FindAllProviders()
	s.providers[1].Feed(airlineSettings["viva"])
}
