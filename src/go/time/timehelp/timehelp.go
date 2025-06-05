package timehelp

import (
	"errors"
	"time"
)

var ErrTimezone = errors.New("invalid location")

func GetLocation(location string) (*time.Location, error) {
	loc, err := time.LoadLocation(location)
	if err != nil {
		return nil, err
	}
	return loc, nil
}

func SetTimeZone(t time.Time, timeZone string) (time.Time, error) {
	loc, err := GetLocation(timeZone)
	if err != nil {
		return time.Time{}, err
	}
	return time.Date(
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(), t.Nanosecond(),
		loc,
	), nil
}
