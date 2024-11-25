package repositories

import (
	"errors"
	"packages/src/amos"
	"packages/src/config"
	"packages/src/pkg/client"
)

type FlightRepository struct{}

func NewFlightRepository() *FlightRepository {
	return &FlightRepository{}
}

func (r *FlightRepository) UpdateAmos(id, password string) error {
	client := client.New(config.Amos.BaseUrl)

	response, err := client.R().
		SetBasicAuth(id, password).
		SetHeader("Content-Type", "application/xml").
		SetBody(amos.NewFutureFlights()).
		Post(config.Amos.FutureFlights.Endpoint)

	if err != nil {
		return err
	}

	if response.StatusCode() != 200 {
		return errors.New(string(response.Body()))
	}

	return nil
}
