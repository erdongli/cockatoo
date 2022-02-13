package conn

import (
	"sync"
	"time"
)

type Stat struct {
	mutex sync.Mutex

	CreatedAt   time.Time
	numSent     int
	numReceived int
}

func newStat() *Stat {
	return &Stat{
		mutex:     sync.Mutex{},
		CreatedAt: time.Now(),
	}
}

func (s *Stat) NumSent() int {
	return s.numSent
}

func (s *Stat) incNumSent() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.numSent++
}

func (s *Stat) NumReceived() int {
	return s.numReceived
}

func (s *Stat) incNumReceived() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.numReceived++
}
