package util

import (
	"strconv"
	"strings"
	"time"
)

func HourToMinute(arg string) (int, error) {
	arr := strings.Split(arg, ":")

	h, err := strconv.Atoi(arr[0])
	if err != nil {
		return 0, err
	}

	m, err := strconv.Atoi(arr[1])
	if err != nil {
		return 0, err
	}

	return h*60 + m, nil
}

func DateFormat(arg *string, from, to string) error {
	date, err := time.Parse(from, *arg)
	if err != nil {
		return err
	}

	*arg = date.Format(to)

	return nil
}
