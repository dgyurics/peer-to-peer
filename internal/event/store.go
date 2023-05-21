package event

import (
	"sync"

	pb "github.com/dgyurics/p2p/api"
)

type Store interface {
	Persist(deviceID, body string) (*pb.Event, error)
	RetrieveAll() []*pb.Event
}

type store struct {
	mu     sync.Mutex
	events []*pb.Event
}

func NewStore() Store {
	return &store{
		mu:     sync.Mutex{},
		events: make([]*pb.Event, 0, 20),
	}
}

func (s *store) Persist(deviceID, body string) (event *pb.Event, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	event, err = New(deviceID, body)
	if err != nil {
		return nil, err
	}
	s.events = append(s.events, event)
	return event, nil
}

func (s *store) RetrieveAll() []*pb.Event {
	return s.events
}
