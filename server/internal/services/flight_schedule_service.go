package services

import (
	"context"
	"packages/server/internal/entities"
	"packages/server/internal/repositories"
	"time"

	"github.com/xuri/excelize/v2"
)

type FlightScheduleService struct {
	_FlightScheduleRepository *repositories.FlightScheduleRepository
}

func NewFlightScheduleService() *FlightScheduleService {
	return &FlightScheduleService{
		_FlightScheduleRepository: repositories.NewFlightScheduleRepository(),
	}
}

func (s *FlightScheduleService) UpdateFlightSchedulesFromWorkbook(context context.Context, workbook *excelize.File) error {
	var flightSchedules []*entities.FlightSchedule

	rows, err := workbook.GetRows(workbook.GetSheetName(0))
	if err != nil {
		return err
	}

	for _, row := range rows {
		if len(row) == 0 {
			continue
		}

		if _, err := time.Parse("02/01/06", row[0]); err == nil {
			if row[4] == "" { // Omit Aircraft Not Assigned Records
				continue
			}

			flightSchedule, err := entities.NewFlightScheduleFromRow(row)
			if err != nil {
				return err
			}

			flightSchedules = append(flightSchedules, flightSchedule)
		}
	}

	for _, flightSchedule := range flightSchedules {
		if err := s._FlightScheduleRepository.UpsertFlightSchedule(context, flightSchedule); err != nil {
			return err
		}
	}

	return nil
}
