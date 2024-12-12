package services

import (
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"packages/src/amos"
	"packages/src/internal/dto"
	"packages/src/internal/entities"
	"packages/src/internal/repositories"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type AmosService struct {
	amosRepository     *repositories.AmosRepository
	currencyRepository *repositories.CurrencyRepository
}

func NewAmosService() *AmosService {
	return &AmosService{
		amosRepository:     repositories.NewAmosRepository(),
		currencyRepository: repositories.NewCurrencyRepository(),
	}
}

func (s *AmosService) GetBasicAuth(context *gin.Context, id, password string) string {
	return base64.StdEncoding.EncodeToString([]byte(id + ":" + password))
}

func (s *AmosService) UpdateCurrency(context *gin.Context, request *dto.UpdateAmosCurrencyRequest, data []entities.ExchangeRate) error {
	if err := s.amosRepository.UpdateCurrency(request.Authorization, data); err != nil {
		return err
	}

	return nil
}

func (s *AmosService) UpdateFutureFlights(context *gin.Context, request *dto.UpdateAmosFutureFlightsRequest) error {
	f, err := excelize.OpenReader(bytes.NewReader(request.Data))
	if err != nil {
		return err
	}

	rows, err := f.GetRows("Sheet")
	if err != nil {
		return err
	}

	var filtered [][]string

	for _, row := range rows {
		if len(row) == 0 {
			continue
		}

		_, err := time.Parse("02/01/06", row[0])
		if err != nil {
			continue
		}

		filtered = append(filtered, row)
	}

	var flights []*amos.FlightsForDateAndAircraft

	for _, row := range filtered {
		flightNumber, err := strconv.Atoi(row[1])
		if err != nil {
			return err
		}

		sdd, err := time.Parse("02/01/06", row[0])
		if err != nil {
			return err
		}

		std, err := time.Parse("02/01/06 15:04", fmt.Sprintf("%s %s", row[0], row[8]))
		if err != nil {
			return err
		}

		sta, err := time.Parse("02/01/06 15:04", fmt.Sprintf("%s %s", row[0], row[9]))
		if err != nil {
			return err
		}

		if sta.Before(std) {
			sta = sta.Add(24 * time.Hour)
		}

		dur := sta.Sub(std).Minutes()
		schedDepartDate := sdd.Format("2006-01-02")
		registration := strings.ReplaceAll(row[4], "HL", "")
		timeFormat := "2006-01-02T15:04:05.00"
		found := false

		for _, flight := range flights {
			if flight.SchedDepartureDate == schedDepartDate && flight.AircraftRegistration == registration {
				found = true

				leg := amos.Leg{
					Carrier:                    row[2],
					FlightNumber:               uint(flightNumber),
					ServiceType:                &row[3],
					ScheduledDepartureDateTime: std.Format(timeFormat),
					ScheduledDepartureAirport:  row[6],
					ScheduledArrivalDateTime:   sta.Format(timeFormat),
					ScheduledArrivalAirport:    row[7],
					EstimatedLegDuration:       uint(dur),
					LegCycles:                  1,
				}

				flight.Legs = append(flight.Legs, leg)
				break
			}
		}

		if !found {
			flights = append(flights, &amos.FlightsForDateAndAircraft{
				SchedDepartureDate:   schedDepartDate,
				AircraftRegistration: registration,
				Legs: []amos.Leg{
					{
						Carrier:                    row[2],
						FlightNumber:               uint(flightNumber),
						ServiceType:                &row[3],
						ScheduledDepartureDateTime: std.Format(timeFormat),
						ScheduledDepartureAirport:  row[6],
						ScheduledArrivalDateTime:   sta.Format(timeFormat),
						ScheduledArrivalAirport:    row[7],
						EstimatedLegDuration:       uint(dur),
						LegCycles:                  1,
					},
				},
			})
		}
	}

	return s.amosRepository.UpdateFutureFlights(request.Authorization, flights)
}

func (s *AmosService) UpdateFutureFlightsWithCSV(context *gin.Context, request *dto.UpdateAmosFutureFlightsRequest) error {
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

	return s.amosRepository.UpdateFutureFlights(request.Authorization, flights)
}
