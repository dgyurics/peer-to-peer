package event

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type Event struct {
	ID       ulid.ULID
	DeviceID string
	Body     string
}

func New(deviceID string, body string) (*Event, error) {
	id, err := ulid.New(ulid.Timestamp(time.Now()), nil)
	if err != nil {
		return nil, err
	}
	return &Event{
		id,
		deviceID,
		body,
	}, nil
}

func (e *Event) Time() time.Time {
	return ulid.Time(e.ID.Time())
}
