package eventuate

import (
	"sync"
	"time"
)

type Signals struct {
	mu       sync.Mutex
	count    int
	ticker   *time.Ticker
	duration time.Duration
	channels map[int]chan int
}

func NewSignals(lcd time.Duration) *Signals {
	t := time.NewTicker(lcd)
	s := &Signals{
		count:    0,
		ticker:   t,
		duration: lcd,
		channels: make(map[int]chan int),
	}

	go func() {
		for range t.C {
			s.mu.Lock()
			s.count++

			for nth, channel := range s.channels {
				if s.count%nth == 0 {
					select {
					case channel <- s.count / nth:
					default:
					}
				}
			}

			s.mu.Unlock()
		}
	}()

	return s
}

func (s *Signals) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.count = 0

	s.ticker.Reset(s.duration)
}

func (s *Signals) Stop() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.ticker.Stop()
}

func (s *Signals) Every(nth int) <-chan int {
	s.mu.Lock()
	defer s.mu.Unlock()

	if channel, exists := s.channels[nth]; exists {
		return channel
	}

	channel := make(chan int, 1)
	s.channels[nth] = channel

	return channel
}
