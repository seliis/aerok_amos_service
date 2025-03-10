package repositories

import (
	"context"
	db "packages/prisma"
	"packages/server/database"
	"packages/server/internal/entities"
)

type FlightScheduleRepository struct{}

func NewFlightScheduleRepository() *FlightScheduleRepository {
	return &FlightScheduleRepository{}
}

func (r *FlightScheduleRepository) UpsertFlightSchedule(context context.Context, entity *entities.FlightSchedule) error {
	if _, err := database.Client.FlightSchedule.UpsertOne(
		db.FlightSchedule.ScheduledDateDepartureFlightNumber(
			db.FlightSchedule.ScheduledDateDeparture.Equals(entity.ScheduledDateDeparture),
			db.FlightSchedule.FlightNumber.Equals(int(entity.FlightNumber)),
		),
	).Create(
		db.FlightSchedule.FlightNumber.Set(int(entity.FlightNumber)),
		db.FlightSchedule.CarrierCode.Set(entity.CarrierCode),
		db.FlightSchedule.ServiceTypeCode.Set(entity.ServiceTypeCode),
		db.FlightSchedule.AircraftRegistration.Set(entity.AircraftRegistration),
		db.FlightSchedule.ScheduledDateDeparture.Set(entity.ScheduledDateDeparture),
		db.FlightSchedule.ScheduledTimeDeparture.Set(entity.ScheduledTimeDeparture),
		db.FlightSchedule.DepartureAirportCode.Set(entity.DepartureAirportCode),
		db.FlightSchedule.ScheduledDateArrival.Set(entity.ScheduledDateArrival),
		db.FlightSchedule.ScheduledTimeArrival.Set(entity.ScheduledTimeArrival),
		db.FlightSchedule.ArrivalAirportCode.Set(entity.ArrivalAirportCode),
		db.FlightSchedule.EstimatedLegDuration.Set(int(entity.EstimatedLegDuration)),
	).Update(
		db.FlightSchedule.CarrierCode.Set(entity.CarrierCode),
		db.FlightSchedule.ServiceTypeCode.Set(entity.ServiceTypeCode),
		db.FlightSchedule.AircraftRegistration.Set(entity.AircraftRegistration),
		db.FlightSchedule.ScheduledTimeDeparture.Set(entity.ScheduledTimeDeparture),
		db.FlightSchedule.DepartureAirportCode.Set(entity.DepartureAirportCode),
		db.FlightSchedule.ScheduledDateArrival.Set(entity.ScheduledDateArrival),
		db.FlightSchedule.ScheduledTimeArrival.Set(entity.ScheduledTimeArrival),
		db.FlightSchedule.ArrivalAirportCode.Set(entity.ArrivalAirportCode),
		db.FlightSchedule.EstimatedLegDuration.Set(int(entity.EstimatedLegDuration)),
	).Exec(context); err != nil {
		return err
	}

	return nil
}

func (r *FlightScheduleRepository) GetFlightSchedulesFromDate(context context.Context, date string) ([]*entities.FlightSchedule, error) {
	models, err := database.Client.FlightSchedule.FindMany(
		db.FlightSchedule.ScheduledDateDeparture.Gte(date),
	).Exec(context)

	if err != nil {
		return nil, err
	}

	var flightSchedules []*entities.FlightSchedule

	for _, model := range models {
		flightSchedules = append(flightSchedules, (*entities.FlightSchedule)(&model.InnerFlightSchedule))
	}

	return flightSchedules, nil
}
