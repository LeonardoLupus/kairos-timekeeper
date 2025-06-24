package interfaces

import (
	"kairos-timekeeper/src/go/time/timebase"
	"kairos-timekeeper/src/go/types"
	"time"
)

type Scheduler interface {
	GetUserID() types.UserID
	AddSlot(t timebase.TimeChatSlot) error
	RemoveSlotAtIndex(index int) error
	AvailabilityAt(t time.Time) timebase.SlotStatus
	AvailabilityAtSlot(t timebase.TimeSpan) timebase.SlotStatus
	GetSlots() []timebase.TimeChatSlot
	SlotCount() int
}

type Participanter interface {
	GetUserID() types.UserID
	GetChatID() types.ChatID
	GetID() (types.UserID, types.ChatID)
	GetUsername() string
	GetTimezone() string

	ChangeTimezone(zoneName string) error
	ChangeUsername(name string)
}
