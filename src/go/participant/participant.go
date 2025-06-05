package participant

import (
	"errors"
	"kairos-timekeeper/src/go/interfaces"
	"kairos-timekeeper/src/go/time/timebase"
	"kairos-timekeeper/src/go/time/timehelp"
	"kairos-timekeeper/src/go/time/timetable"
)

var ErrChangeTimezone = errors.New("invalid location to try change timezone")

type participant struct {
	ID       int64
	Username string
	Timezone string
	Schedule interfaces.Scheduler
}

func NewParticipant(id int64, name, location string) (interfaces.Participanter, error) {
	_, err := timehelp.GetLocation(location)
	if err != nil {
		return nil, err
	}
	p := participant{
		ID:       id,
		Username: name,
		Timezone: location,
		Schedule: timetable.NewTimeSheet(),
	}
	return &p, nil
}

func (p *participant) GetID() int64 {
	return p.ID
}

func (p *participant) GetUsername() string {
	return p.Username
}

func (p *participant) GetTimezone() string {
	return p.Timezone
}

func (p *participant) GetSchedule() interfaces.Scheduler {
	return p.Schedule
}

func (p *participant) AddTimeSlot(t timebase.TimeSlot) error {
	tStart, err0 := timehelp.SetTimeZone(t.Start, p.Timezone)
	tEnd, err1 := timehelp.SetTimeZone(t.End, p.Timezone)

	if err0 != nil || err1 != nil {
		return err0
	}

	tStart = tStart.UTC()
	tEnd = tEnd.UTC()

	return p.Schedule.AddSlot(timebase.TimeSlot{
		Status: t.Status,
		TimeSpan: timebase.TimeSpan{
			Start: tStart,
			End:   tEnd,
		},
	})
}

func (p *participant) RemoveTimeSlotAtIndex(index int) error {
	return p.Schedule.RemoveSlotAtIndex(index)
}

func (p *participant) ChangeTimezone(zoneName string) error {
	_, err := timehelp.GetLocation(zoneName)
	if err != nil {
		return ErrChangeTimezone
	}
	p.Timezone = zoneName
	return nil
}
