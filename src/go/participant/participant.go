package participant

import (
	"errors"
	"kairos-timekeeper/src/go/interfaces"
	"kairos-timekeeper/src/go/time/timehelp"
	"kairos-timekeeper/src/go/types"
)

var ErrChangeTimezone = errors.New("invalid location to try change timezone")

type participant struct {
	UserID   types.UserID
	ChatID   types.ChatID
	Username string
	Timezone string
}

func NewParticipant(userId types.UserID, chatId types.ChatID, name, location string) (interfaces.Participanter, error) {
	_, err := timehelp.GetLocation(location)
	if err != nil {
		return nil, err
	}
	p := participant{
		UserID:   userId,
		ChatID:   chatId,
		Username: name,
		Timezone: location,
	}
	return &p, nil
}

func (p *participant) GetUserID() types.UserID {
	return p.UserID
}

func (p *participant) GetChatID() types.ChatID {
	return p.ChatID
}

func (p *participant) GetID() (types.UserID, types.ChatID) {
	return p.GetUserID(), p.GetChatID()
}

func (p *participant) GetUsername() string {
	return p.Username
}

func (p *participant) GetTimezone() string {
	return p.Timezone
}

func (p *participant) ChangeTimezone(zoneName string) error {
	_, err := timehelp.GetLocation(zoneName)
	if err != nil {
		return ErrChangeTimezone
	}
	p.Timezone = zoneName
	return nil
}

func (p *participant) ChangeUsername(name string) {
	p.Username = name
}
