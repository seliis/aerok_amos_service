package services

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"packages/src/amos"
	"packages/src/internal/dto"
	"packages/src/internal/repositories"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type FlightService struct {
	flightRepository *repositories.FlightRepository
}

func NewFlightService() *FlightService {
	return &FlightService{flightRepository: repositories.NewFlightRepository()}
}

func (s *FlightService) UpdateAmos(context *gin.Context, request *dto.UpdateAmosFutureFlightsRequest) error {
	if err := amos.CheckWebServiceAuth(request.WebServiceAuth); err != nil {
		return err
	}

	data := csv.NewReader(bytes.NewReader(request.Data))

	records, err := data.ReadAll()
	if err != nil {
		return err
	}

	var filtered [][]string

	for _, record := range records {
		if len(record) == 0 {
			continue
		}

		_, err := time.Parse("2006-01-02", record[0])
		if err != nil {
			continue
		}

		if len(record[4]) == 0 {
			continue
		}

		filtered = append(filtered, record)
	}

	var flights []*amos.FlightsForDateAndAircraft

	for _, record := range filtered {
		flightNumber, err := strconv.Atoi(record[1])
		if err != nil {
			return err
		}

		std, err := time.Parse("2006-01-02 15:04", fmt.Sprintf("%s %s", record[0], record[7]))
		if err != nil {
			return err
		}

		sta, err := time.Parse("2006-01-02 15:04", fmt.Sprintf("%s %s", record[0], record[8]))
		if err != nil {
			return err
		}

		if sta.Before(std) {
			sta = sta.Add(24 * time.Hour)
		}

		dur := sta.Sub(std).Minutes()

		registration := strings.ReplaceAll(record[4], "HL", "")
		timeFormat := "2006-01-02T15:04:05.00"
		found := false

		for _, flight := range flights {
			if flight.SchedDepartureDate == record[0] && flight.AircraftRegistration == registration {
				found = true

				leg := amos.Leg{
					Carrier:                    record[2],
					FlightNumber:               uint(flightNumber),
					ScheduledDepartureDateTime: std.Format(timeFormat),
					ScheduledDepartureAirport:  record[5],
					ScheduledArrivalDateTime:   sta.Format(timeFormat),
					ScheduledArrivalAirport:    record[6],
					EstimatedLegDuration:       uint(dur),
					LegCycles:                  1,
				}

				flight.Legs = append(flight.Legs, leg)
				break
			}
		}

		if !found {
			flights = append(flights, &amos.FlightsForDateAndAircraft{
				SchedDepartureDate:   record[0],
				AircraftRegistration: registration,
				Legs: []amos.Leg{
					{
						Carrier:                    record[2],
						FlightNumber:               uint(flightNumber),
						ScheduledDepartureDateTime: std.Format(timeFormat),
						ScheduledDepartureAirport:  record[5],
						ScheduledArrivalDateTime:   sta.Format(timeFormat),
						ScheduledArrivalAirport:    record[6],
						EstimatedLegDuration:       uint(dur),
						LegCycles:                  1,
					},
				},
			})
		}
	}

	return s.flightRepository.UpdateAmos(request.WebServiceAuth.Id, request.WebServiceAuth.Password, flights)
}
