package event

// not best place to put this, but good enough for proof of concept app

type Store interface {
	Persist(deviceID, body string) (*Event, error)
	RetrieveAll() []Event
}

type store struct {
	events []Event
}

func NewStore() Store {
	return &store{
		events: make([]Event, 0, 20), // arbitrary capacity, just to avoid reallocs
	}
}

func (s *store) Persist(deviceID, body string) (event *Event, err error) {
	event, err = New(deviceID, body)
	if err != nil {
		return nil, err
	}
	s.events = append(s.events, *event)
	return event, nil
}

func (s *store) RetrieveAll() []Event {
	return s.events
}
