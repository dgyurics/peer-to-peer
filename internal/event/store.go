package event

import (
	"sync"

	pb "github.com/dgyurics/p2p/api"
)

type Store interface {
	Persist(event *pb.Event)
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

func (s *store) Persist(event *pb.Event) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events = append(s.events, event)
}

func (s *store) RetrieveAll() []*pb.Event {
	return s.events
}
