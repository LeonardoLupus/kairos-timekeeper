package participant

import (
	"errors"
	"kairos-timekeeper/src/go/timehelp"
	"kairos-timekeeper/src/go/timetable"
)

var ErrChangeTimezone = errors.New("invalid location to try change timezone")

type Participanter interface {
	GetID() int64
	GetUsername() string
	GetTimezone() string
	GetSchedule() timetable.TimeSheeter

	AddTimeSlot(t timetable.TimeSlot) error
	RemoveTimeSlotAtIndex(index int) error
	ChangeTimezone(zoneName string) error
}

type participant struct {
	ID       int64
	Username string
	Timezone string
	Schedule timetable.TimeSheeter
}

func NewParticipant(id int64, name, location string) (Participanter, error) {
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

func (p *participant) GetSchedule() timetable.TimeSheeter {
	return p.Schedule
}

func (p *participant) AddTimeSlot(t timetable.TimeSlot) error {
	tStart, err0 := timehelp.SetTimeZone(t.Start, p.Timezone)
	tEnd, err1 := timehelp.SetTimeZone(t.End, p.Timezone)

	if err0 != nil || err1 != nil {
		return err0
	}

	tStart = tStart.UTC()
	tEnd = tEnd.UTC()

	return p.Schedule.AddSlot(timetable.TimeSlot{
		Status: t.Status,
		TimeSpan: timetable.TimeSpan{
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
