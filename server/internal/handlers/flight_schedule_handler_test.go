package handlers_test

import (
	"os"
	"testing"

	"resty.dev/v3"
)

func TestFlightScheduleHandler(t *testing.T) {
	c := resty.New()

	f, err := os.Open("../../../data/aims.xlsx")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	req := c.R().SetFileReader("flight_schedule", "aims.xlsx", f)

	res, err := req.Post("http://127.0.0.1:8080/api/flight-schedule/update")
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode() != 200 {
		t.Errorf("err: %d", res.StatusCode())
	}
}
