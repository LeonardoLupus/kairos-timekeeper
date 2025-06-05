package interfaces

import (
	"kairos-timekeeper/src/go/time/timebase"
	"time"
)

type Scheduler interface {
	AddSlot(t timebase.TimeSlot) error
	RemoveSlotAtIndex(index int) error
	AvailabilityAt(t time.Time) timebase.SlotStatus
	GetSlots() []timebase.TimeSlot
	SlotCount() int
	AvailabilityAtSlot(t timebase.TimeSpan) timebase.SlotStatus
}

type Participanter interface {
	GetID() int64
	GetUsername() string
	GetTimezone() string

	AddTimeSlot(t timebase.TimeSlot) error
	RemoveTimeSlotAtIndex(index int) error
	ChangeTimezone(zoneName string) error
}
