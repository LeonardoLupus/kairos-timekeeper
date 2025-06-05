package timebase

import (
	"errors"
	"time"
)

var ErrInvalidTimeInterval = errors.New("invalid time interval (end before start)")
var ErrInvalidTimeIntersect = errors.New("invalid time interval (intervals is intersect)")
var ErrInvalidSlotIndex = errors.New("invalid slot index")

type SlotStatus int

const (
	Preferred SlotStatus = iota
	Neutral
	Unavailable
)

type TimeSpan struct {
	Start time.Time
	End   time.Time
}

type TimeSlot struct {
	TimeSpan
	Status SlotStatus
}

func (tsp *TimeSpan) Contains(t time.Time) bool {
	return t.Equal(tsp.Start) || (t.After(tsp.Start) && t.Before(tsp.End))
}

func (tsp TimeSpan) Overlaps(other TimeSpan) bool {
	return tsp.Start.Before(other.End) && tsp.End.After(other.Start)
}

func (tsp *TimeSpan) ChangeStart(t time.Time) error {
	if t.Before(tsp.End) {
		tsp.Start = t
		return nil
	}
	return ErrInvalidTimeInterval
}

func (tsp *TimeSpan) ChangeEnd(t time.Time) error {
	if tsp.Start.Before(t) {
		tsp.End = t
		return nil
	}
	return ErrInvalidTimeInterval
}

func (tsp *TimeSpan) ChangeSpan(tStart, tEnd time.Time) error {
	if tStart.Before(tEnd) {
		tsp.Start = tStart
		tsp.End = tEnd
		return nil
	}
	return ErrInvalidTimeInterval
}

func (ts *TimeSlot) ChangeStatus(s SlotStatus) {
	ts.Status = s
}
