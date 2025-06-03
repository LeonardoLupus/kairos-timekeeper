package timetable

import (
	"errors"
	"slices"
	"time"
)

type SlotStatus int

var ErrInvalidTimeInterval = errors.New("invalid time interval (end before start)")
var ErrInvalidTimeIntersect = errors.New("invalid time interval (intervals is intersect)")
var ErrInvalidSlotIndex = errors.New("invalid slot index")

const (
	Neutral SlotStatus = iota
	Preferred
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

type TimeSheet struct {
	Slots []TimeSlot
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

func (ts *TimeSheet) sortSlots() {
	slices.SortFunc(ts.Slots, func(a, b TimeSlot) int {
		return a.Start.Compare(b.Start)
	})
}

func (ts *TimeSheet) IsIntersect(span TimeSpan) bool {
	for _, v := range ts.Slots {
		if v.TimeSpan.Overlaps(span) {
			return true
		}
	}
	return false
}

func (ts *TimeSheet) AddSlot(t TimeSlot) error {
	if !t.End.After(t.Start) {
		return ErrInvalidTimeInterval
	}
	if ts.IsIntersect(t.TimeSpan) {
		return ErrInvalidTimeIntersect
	}
	ts.Slots = append(ts.Slots, t)
	ts.sortSlots()
	return nil
}

func (ts *TimeSheet) DelSlotAtIndex(index int) error {
	if index < 0 || index >= len(ts.Slots) {
		return ErrInvalidSlotIndex
	}
	ts.Slots = slices.Delete(ts.Slots, index, index+1)
	return nil
}

func (ts *TimeSheet) AvailabilityAt(t time.Time) SlotStatus {
	for _, v := range ts.Slots {
		if v.Contains(t) {
			return v.Status
		}
	}
	return Neutral
}

func (ts *TimeSheet) FindSlotsByStatus(status SlotStatus) []TimeSlot {
	slots := make([]TimeSlot, 0, len(ts.Slots))
	for _, v := range ts.Slots {
		if v.Status == status {
			slots = append(slots, v)
		}
	}
	return slots
}
