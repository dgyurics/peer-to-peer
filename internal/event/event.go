package event

import (
	"time"

	pb "github.com/dgyurics/p2p/api"
	"github.com/oklog/ulid/v2"
)

func New(deviceID string, body string) (*pb.Event, error) {
	id, err := ulid.New(ulid.Timestamp(time.Now()), nil)
	if err != nil {
		return nil, err
	}
	return &pb.Event{
		Id:       id.String(),
		DeviceId: deviceID,
		Body:     body,
	}, nil
}

func Time(id string) (time.Time, error) {
	e, err := ulid.ParseStrict(id)
	if err != nil {
		return time.Time{}, err
	}
	return ulid.Time(e.Time()), nil
}
