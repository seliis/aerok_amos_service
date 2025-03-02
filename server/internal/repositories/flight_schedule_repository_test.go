package repositories_test

import (
	"context"
	"packages/server/database"
	"packages/server/internal/entities"
	"packages/server/internal/repositories"
	"testing"
)

func TestFlightScheduleRepository(t *testing.T) {
	if err := database.Connect("file:../../../prisma/database.db"); err != nil {
		t.Error(err)
	}

	r := repositories.NewFlightScheduleRepository()

	flightSchedules := []*entities.FlightSchedule{
		{
			FlightNumber:           322,
			CarrierCode:            "RF",
			ServiceTypeCode:        "J",
			AircraftRegistration:   "HL8562",
			ScheduledDateDeparture: "2025-03-01",
			ScheduledTimeDeparture: "00:45",
			DepartureAirportCode:   "CJJ",
			ScheduledDateArrival:   "2025-03-01",
			ScheduledTimeArrival:   "02:55",
			ArrivalAirportCode:     "NRT",
		},
		{
			FlightNumber:           321,
			CarrierCode:            "RF",
			ServiceTypeCode:        "J",
			AircraftRegistration:   "HL8562",
			ScheduledDateDeparture: "2025-03-01",
			ScheduledTimeDeparture: "03:55",
			DepartureAirportCode:   "NRT",
			ScheduledDateArrival:   "2025-03-01",
			ScheduledTimeArrival:   "06:25",
			ArrivalAirportCode:     "CJJ",
		},
	}

	t.Run("UpsertFlightSchedules", func(t *testing.T) {
		for _, flightSchedule := range flightSchedules {
			if err := r.UpsertFlightSchedule(context.Background(), flightSchedule); err != nil {
				t.Error(err)
			}
		}
	})

	defer func() {
		if err := database.Disconnect(); err != nil {
			t.Error(err)
		}
	}()
}
