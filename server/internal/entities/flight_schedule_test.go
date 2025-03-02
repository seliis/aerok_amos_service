package entities_test

import (
	"encoding/json"
	"fmt"
	"packages/server/internal/entities"
	"testing"
	"time"

	"github.com/xuri/excelize/v2"
)

func TestFlightSchedule(t *testing.T) {
	f, err := excelize.OpenFile("../../../data/aims.xlsx")
	if err != nil {
		t.Error(err)
	}

	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		t.Error(err)
	}

	var data []*entities.FlightSchedule

	for _, row := range rows {
		if len(row) == 0 {
			continue
		}

		if _, err := time.Parse("02/01/06", row[0]); err == nil {
			entity, err := entities.NewFlightScheduleFromRow(row)
			if err != nil {
				t.Error(err)
			}

			data = append(data, entity)
		}
	}

	res, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%s\n", res)

	defer func() {
		if err := f.Close(); err != nil {
			t.Error(err)
		}
	}()
}
