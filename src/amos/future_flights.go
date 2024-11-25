package amos

import (
	"encoding/xml"
)

type _FutureFlights struct {
	XMLName                    string                       `xml:"future_flights"`
	PublishedAt                xml.Attr                     `xml:"published_at,attr"`
	Version                    xml.Attr                     `xml:"version,attr"`
	FlightsForDateAndAircrafts []_FlightsForDateAndAircraft `xml:"flights_for_date_and_aircraft"`
}

type _FlightsForDateAndAircraft struct {
	SchedDepartureDate   string `xml:"sched_departure_date"`
	AircraftRegistration string `xml:"aircraft_registration"`
	Legs                 []_Leg `xml:"leg"`
}

type _Leg struct {
	Carrier                                  string        `xml:"carrier"`
	FlightNumber                             uint          `xml:"flightNumber"`
	FlightSuffix                             string        `xml:"flightSuffix"` // Optional
	ServiceType                              string        `xml:"serviceType"`  // Optional
	ScheduledDepartureDateTime               string        `xml:"scheduledDepartureDateTime"`
	EstimatedDepartureDateTime               string        `xml:"estimatedDepartureDateTime"` // Optional
	ScheduledDepartureAirport                string        `xml:"scheduledDepartureAirport"`
	DepartureStand                           string        `xml:"departureStand"` // Optional
	DepartureGate                            string        `xml:"departureGate"`  // Optional
	ScheduledArrivalDateTime                 string        `xml:"scheduledArrivalDateTime"`
	EstimatedArrivalDateTime                 string        `xml:"estimatedArrivalDateTime"` // Optional
	ScheduledArrivalAirport                  string        `xml:"scheduledArrivalAirport"`
	ArrivalStand                             string        `xml:"arrivalStand"` // Optional
	ArrivalGate                              string        `xml:"arrivalGate"`  // Optional
	EstimatedLegDuration                     uint          `xml:"estimatedLegDuration"`
	LegCycles                                uint          `xml:"legCycles"`
	EstimatedTotalAircraftMinutesAtTouchdown uint          `xml:"estimatedTotalAircraftMinutesAtTouchdown"` // Optional
	EstimatedTotalAircraftCyclesAtTouchdown  uint          `xml:"estimatedTotalAircraftCyclesAtTouchdown"`  // Optional
	ExternalSystemLegId                      uint          `xml:"externalSystemLegId"`                      // Optional
	Stickers                                 []_Sticker    `xml:"sticker"`                                  // Optional
	CrewMembers                              []_CrewMember `xml:"crewMember"`                               // Optional
	State                                    string        `xml:"state"`                                    // Optional
}

type _Sticker struct {
	Header      string `xml:"header"`
	StickerText string `xml:"stickerText"`
}

type _CrewMember struct {
	LetterCode   string `xml:"letterCode"`
	FunctionCode string `xml:"functionCode"`
}

func NewFutureFlights() _FutureFlights {
	testData := []_FlightsForDateAndAircraft{
		{
			SchedDepartureDate:   "2024-12-01",
			AircraftRegistration: "HL8385",
			Legs: []_Leg{
				{
					Carrier:                                  "EOK",
					FlightNumber:                             321,
					FlightSuffix:                             "K",
					ServiceType:                              "J", // IATA Code; J: Scheduled Flight, Passenger, Normal Service
					ScheduledDepartureDateTime:               "2024-12-01T00:00:00.00",
					EstimatedDepartureDateTime:               "2024-12-01T00:10:00.00",
					ScheduledDepartureAirport:                "CJJ",
					DepartureStand:                           "314",
					DepartureGate:                            "314",
					ScheduledArrivalDateTime:                 "2024-12-01T02:00:00.00",
					EstimatedArrivalDateTime:                 "2024-12-01T02:10:00.00",
					ScheduledArrivalAirport:                  "NRT",
					ArrivalStand:                             "419",
					ArrivalGate:                              "419",
					EstimatedLegDuration:                     120,
					LegCycles:                                1,
					EstimatedTotalAircraftMinutesAtTouchdown: 120,
					EstimatedTotalAircraftCyclesAtTouchdown:  1,
					ExternalSystemLegId:                      1,
					Stickers: []_Sticker{
						{
							Header:      "Header",
							StickerText: "StickerText",
						},
					},
					CrewMembers: []_CrewMember{
						{
							LetterCode:   "123456",
							FunctionCode: "CPT",
						},
					},
					State: "N",
				},
			},
		},
	}

	return _FutureFlights{
		PublishedAt: xml.Attr{
			Name:  xml.Name{Local: "published_at"},
			Value: "2024-11-22T00:00:00.00",
		},
		Version: xml.Attr{
			Name:  xml.Name{Local: "version"},
			Value: "1.1",
		},
		FlightsForDateAndAircrafts: testData,
	}
}
