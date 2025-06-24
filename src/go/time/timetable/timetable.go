package timetable

import (
	"kairos-timekeeper/src/go/interfaces"
	"kairos-timekeeper/src/go/time/timebase"
	"kairos-timekeeper/src/go/types"
	"slices"
	"time"
)

type TimeSheet struct {
	UserID types.UserID
	Slots  []timebase.TimeChatSlot
}

func NewTimeSheet(user types.UserID, slots []timebase.TimeChatSlot) interfaces.Scheduler {
	return &TimeSheet{
		UserID: user,
		Slots:  slots,
	}
}

func (ts *TimeSheet) sortSlots() {
	slices.SortFunc(ts.Slots, func(a, b timebase.TimeChatSlot) int {
		return a.Start.Compare(b.Start)
	})
}

func (ts *TimeSheet) IsIntersect(span timebase.TimeSpan) bool {
	for _, v := range ts.Slots {
		if v.TimeSpan.Overlaps(span) {
			return true
		}
	}
	return false
}

func (ts *TimeSheet) GetUserID() types.UserID {
	return ts.UserID
}

func (ts *TimeSheet) AddSlot(t timebase.TimeChatSlot) error {
	if !t.End.After(t.Start) {
		return timebase.ErrInvalidTimeInterval
	}
	if ts.IsIntersect(t.TimeSpan) {
		return timebase.ErrInvalidTimeIntersect
	}
	ts.Slots = append(ts.Slots, t)
	ts.sortSlots()
	return nil
}

func (ts *TimeSheet) RemoveSlotAtIndex(index int) error {
	if index < 0 || index >= len(ts.Slots) {
		return timebase.ErrInvalidSlotIndex
	}
	ts.Slots = slices.Delete(ts.Slots, index, index+1)
	return nil
}

func (ts *TimeSheet) FindSlotsByStatus(status timebase.SlotStatus) []timebase.TimeChatSlot {
	slots := make([]timebase.TimeChatSlot, 0, len(ts.Slots))
	for _, v := range ts.Slots {
		if v.Status == status {
			slots = append(slots, v)
		}
	}
	return slots
}

func (ts *TimeSheet) AvailabilityAt(t time.Time) timebase.SlotStatus {
	for _, v := range ts.Slots {
		if v.Contains(t) {
			return v.Status
		}
	}
	return timebase.Neutral
}

func (ts *TimeSheet) AvailabilityAtSlot(t timebase.TimeSpan) timebase.SlotStatus {
	status := timebase.Preferred
	for _, v := range ts.Slots {
		if v.Overlaps(t) {
			if v.Status > status {
				status = v.Status
			}
		}
	}
	return status
}

func (ts *TimeSheet) GetSlots() []timebase.TimeChatSlot {
	return slices.Clone(ts.Slots)
}

func (ts *TimeSheet) SlotCount() int {
	return len(ts.Slots)
}
