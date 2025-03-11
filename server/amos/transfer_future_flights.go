package amos

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"packages/server/client"
	"packages/server/config"
	"packages/server/internal/entities"
	"time"
)

type _FutureFlights struct {
	XMLName                    string                        `xml:"future_flights"`
	PublishedAt                xml.Attr                      `xml:"published_at,attr"`
	Version                    xml.Attr                      `xml:"version,attr"`
	FlightsForDateAndAircrafts []*_FlightsForDateAndAircraft `xml:"flights_for_date_and_aircraft"`
}

type _FlightsForDateAndAircraft struct {
	SchedDepartureDate   string  `xml:"sched_departure_date"`
	AircraftRegistration string  `xml:"aircraft_registration"`
	Legs                 []*_Leg `xml:"leg"`
}

type _Leg struct {
	Carrier                                  string        `xml:"carrier"`
	FlightNumber                             uint          `xml:"flightNumber"`
	FlightSuffix                             *string       `xml:"flightSuffix"` // Optional
	ServiceType                              *string       `xml:"serviceType"`  // Optional
	ScheduledDepartureDateTime               string        `xml:"scheduledDepartureDateTime"`
	EstimatedDepartureDateTime               *string       `xml:"estimatedDepartureDateTime"` // Optional
	ScheduledDepartureAirport                string        `xml:"scheduledDepartureAirport"`
	DepartureStand                           *string       `xml:"departureStand"` // Optional
	DepartureGate                            *string       `xml:"departureGate"`  // Optional
	ScheduledArrivalDateTime                 string        `xml:"scheduledArrivalDateTime"`
	EstimatedArrivalDateTime                 *string       `xml:"estimatedArrivalDateTime"` // Optional
	ScheduledArrivalAirport                  string        `xml:"scheduledArrivalAirport"`
	ArrivalStand                             *string       `xml:"arrivalStand"` // Optional
	ArrivalGate                              *string       `xml:"arrivalGate"`  // Optional
	EstimatedLegDuration                     uint          `xml:"estimatedLegDuration"`
	LegCycles                                uint          `xml:"legCycles"`
	EstimatedTotalAircraftMinutesAtTouchdown *uint         `xml:"estimatedTotalAircraftMinutesAtTouchdown"` // Optional
	EstimatedTotalAircraftCyclesAtTouchdown  *uint         `xml:"estimatedTotalAircraftCyclesAtTouchdown"`  // Optional
	ExternalSystemLegId                      *uint         `xml:"externalSystemLegId"`                      // Optional
	Stickers                                 []_Sticker    `xml:"sticker"`                                  // Optional
	CrewMembers                              []_CrewMember `xml:"crewMember"`                               // Optional
	State                                    *string       `xml:"state"`                                    // Optional
}

type _Sticker struct {
	Header      string `xml:"header"`
	StickerText string `xml:"stickerText"`
}

type _CrewMember struct {
	LetterCode   string `xml:"letterCode"`
	FunctionCode string `xml:"functionCode"`
}

func NewFutureFlights(flightSchedules []*entities.FlightSchedule) _FutureFlights {
	// testData := []_FlightsForDateAndAircraft{
	// 	{
	// 		SchedDepartureDate:   "2024-12-01",
	// 		AircraftRegistration: "8385",
	// 		Legs: []*_Leg{
	// 			{
	// 				Carrier:                                  "RF",
	// 				FlightNumber:                             321,
	// 				FlightSuffix:                             nil,
	// 				ServiceType:                              nil, // IATA Code; J: Scheduled Flight, Passenger, Normal Service
	// 				ScheduledDepartureDateTime:               "2024-12-01T00:00:00.00",
	// 				EstimatedDepartureDateTime:               nil,
	// 				ScheduledDepartureAirport:                "CJJ",
	// 				DepartureStand:                           nil,
	// 				DepartureGate:                            nil,
	// 				ScheduledArrivalDateTime:                 "2024-12-01T02:00:00.00",
	// 				EstimatedArrivalDateTime:                 nil,
	// 				ScheduledArrivalAirport:                  "NRT",
	// 				ArrivalStand:                             nil,
	// 				ArrivalGate:                              nil,
	// 				EstimatedLegDuration:                     120,
	// 				LegCycles:                                1,
	// 				EstimatedTotalAircraftMinutesAtTouchdown: nil,
	// 				EstimatedTotalAircraftCyclesAtTouchdown:  nil,
	// 				ExternalSystemLegId:                      nil,
	// 				Stickers:                                 nil,
	// 				CrewMembers:                              nil,
	// 				State:                                    nil,
	// 			},
	// 		},
	// 	},
	// }

	var flightsForDateAndAircrafts []*_FlightsForDateAndAircraft

	isExist := func(schedDepartureDate string, aircraftRegistration string) bool {
		for _, flightsForDateAndAircraft := range flightsForDateAndAircrafts {
			if flightsForDateAndAircraft.SchedDepartureDate == schedDepartureDate && flightsForDateAndAircraft.AircraftRegistration == aircraftRegistration {
				return true
			}
		}

		return false
	}

	for _, flightSchedule := range flightSchedules {
		if !isExist(flightSchedule.ScheduledDateDeparture, flightSchedule.AircraftRegistration) {
			flightsForDateAndAircrafts = append(flightsForDateAndAircrafts, &_FlightsForDateAndAircraft{
				SchedDepartureDate:   flightSchedule.ScheduledDateDeparture,
				AircraftRegistration: flightSchedule.AircraftRegistration,
			})
		}
	}

	for _, flightSchedule := range flightSchedules {
		for _, flightForDateAndAircraft := range flightsForDateAndAircrafts {
			if flightForDateAndAircraft.SchedDepartureDate == flightSchedule.ScheduledDateDeparture && flightForDateAndAircraft.AircraftRegistration == flightSchedule.AircraftRegistration {
				flightForDateAndAircraft.Legs = append(flightForDateAndAircraft.Legs, &_Leg{
					Carrier:      flightSchedule.CarrierCode,
					FlightNumber: uint(flightSchedule.FlightNumber),
					FlightSuffix: func() *string {
						if flightSchedule.FlightSuffix == "" {
							return nil
						}

						return &flightSchedule.FlightSuffix
					}(),
					ServiceType: func() *string {
						if flightSchedule.ServiceTypeCode == "" {
							return nil
						}

						return &flightSchedule.ServiceTypeCode
					}(),
					ScheduledDepartureDateTime:               fmt.Sprintf("%sT%s:00.00", flightSchedule.ScheduledDateDeparture, flightSchedule.ScheduledTimeDeparture),
					EstimatedDepartureDateTime:               nil,
					ScheduledDepartureAirport:                flightSchedule.DepartureAirportCode,
					DepartureStand:                           nil,
					DepartureGate:                            nil,
					ScheduledArrivalDateTime:                 fmt.Sprintf("%sT%s:00.00", flightSchedule.ScheduledDateArrival, flightSchedule.ScheduledTimeArrival),
					EstimatedArrivalDateTime:                 nil,
					ScheduledArrivalAirport:                  flightSchedule.ArrivalAirportCode,
					ArrivalStand:                             nil,
					ArrivalGate:                              nil,
					EstimatedLegDuration:                     uint(flightSchedule.EstimatedLegDuration),
					LegCycles:                                1,
					EstimatedTotalAircraftMinutesAtTouchdown: nil,
					EstimatedTotalAircraftCyclesAtTouchdown:  nil,
					ExternalSystemLegId:                      nil,
					Stickers:                                 nil,
					CrewMembers:                              nil,
					State:                                    nil,
				})
			}
		}
	}

	return _FutureFlights{
		PublishedAt: xml.Attr{
			Name: xml.Name{
				Local: "published_at",
			},
			Value: time.Now().Format("2006-01-02T15:04:05.00Z"),
		},
		Version: xml.Attr{
			Name: xml.Name{
				Local: "version",
			},
			Value: "1.1",
		},
		FlightsForDateAndAircrafts: flightsForDateAndAircrafts,
	}
}

func (futureFlights *_FutureFlights) Push() error {
	config := config.AMOS.Services.TransferFutureFlights

	r, err := client.GetAmosRequest().SetBody(futureFlights).Post(config.EndPoint)
	if err != nil {
		return err
	}

	if r.StatusCode() != http.StatusOK {
		return errors.New("transfer future flights failed")
	}

	return nil
}
