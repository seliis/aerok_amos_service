package entities

import (
	"errors"
	"packages/server/util"
	"strconv"
	"strings"
	"time"
)

type FlightSchedule struct {
	ID                     string `json:"id"`
	FlightNumber           int    `json:"flight_number"`
	FlightSuffix           string `json:"flight_suffix"`
	CarrierCode            string `json:"carrier_code"`
	ServiceTypeCode        string `json:"service_type_code"`
	AircraftRegistration   string `json:"aircraft_registration"`
	ScheduledDateDeparture string `json:"scheduled_date_departure"`
	ScheduledTimeDeparture string `json:"scheduled_time_departure"`
	DepartureAirportCode   string `json:"departure_airport_code"`
	ScheduledDateArrival   string `json:"scheduled_date_arrival"`
	ScheduledTimeArrival   string `json:"scheduled_time_arrival"`
	ArrivalAirportCode     string `json:"arrival_airport_code"`
	EstimatedLegDuration   int    `json:"estimated_leg_duration"`
}

func NewFlightScheduleFromRow(row []string) (*FlightSchedule, error) {
	entity := &FlightSchedule{
		FlightNumber:           0,
		FlightSuffix:           "",
		CarrierCode:            row[2],
		ServiceTypeCode:        row[3],
		AircraftRegistration:   row[4],
		ScheduledDateDeparture: row[0],
		ScheduledTimeDeparture: row[8],
		DepartureAirportCode:   row[6],
		ScheduledDateArrival:   row[0],
		ScheduledTimeArrival:   row[9],
		ArrivalAirportCode:     row[7],
		EstimatedLegDuration:   0,
	}

	if flightNumber, flightSuffix, err := func(value string) (int, string, error) {
		var numeric string
		var suffix string
		numeric = value

		for i, c := range value {
			if c < '0' || c > '9' {
				numeric = value[:i]
				suffix = value[i:]
				break
			}
		}

		converted, err := strconv.Atoi(numeric)
		if err != nil {
			return 0, "", err
		}

		return converted, suffix, nil
	}(row[1]); err != nil {
		return nil, err
	} else {
		entity.FlightNumber = flightNumber
		entity.FlightSuffix = flightSuffix
	}

	entity.AircraftRegistration = strings.ReplaceAll(entity.AircraftRegistration, "HL", "")

	if err := util.DateFormat(&entity.ScheduledDateDeparture, "02/01/06", "2006-01-02"); err != nil {
		return nil, err
	}

	if err := util.DateFormat(&entity.ScheduledDateArrival, "02/01/06", "2006-01-02"); err != nil {
		return nil, err
	}

	std, err := util.HourToMinute(entity.ScheduledTimeDeparture)
	if err != nil {
		return nil, err
	}

	sta, err := util.HourToMinute(entity.ScheduledTimeArrival)
	if err != nil {
		return nil, err
	}

	if std > sta {
		date, err := time.Parse("2006-01-02", entity.ScheduledDateDeparture)
		if err != nil {
			return nil, err
		}

		entity.ScheduledDateArrival = date.AddDate(0, 0, 1).Format("2006-01-02")
	}

	if std > sta {
		entity.EstimatedLegDuration = (sta + (24 * 60)) - std
	} else {
		entity.EstimatedLegDuration = sta - std
	}

	if entity.EstimatedLegDuration == 0 {
		return nil, errors.New("estimated leg duration is zero")
	}

	return entity, nil
}
