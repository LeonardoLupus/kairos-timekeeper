package interfaces

import (
	"kairos-timekeeper/src/go/time/timebase"
	"kairos-timekeeper/src/go/types"
	"time"
)

type Scheduler interface {
	GetUserID() types.UserID
	GetChatID() types.ChatID
	GetID() (types.UserID, types.ChatID)
	AddSlot(t timebase.TimeSlot) error
	RemoveSlotAtIndex(index int) error
	AvailabilityAt(t time.Time) timebase.SlotStatus
	GetSlots() []timebase.TimeSlot
	SlotCount() int
	AvailabilityAtSlot(t timebase.TimeSpan) timebase.SlotStatus
}

type Participanter interface {
	GetUserID() types.UserID
	GetChatID() types.ChatID
	GetID() (types.UserID, types.ChatID)
	GetUsername() string
	GetTimezone() string

	AddTimeSlot(t timebase.TimeSlot) error
	RemoveTimeSlotAtIndex(index int) error
	ChangeTimezone(zoneName string) error
	ChangeUsername(name string)
}
