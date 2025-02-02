package eventuate

import (
	"sync"
	"time"
)

type Countdown struct {
	mu       sync.Mutex
	start    time.Time
	end      time.Time
	timer    *time.Timer
	duration time.Duration
}

func NewCountdown(duration time.Duration) *Countdown {
	start := time.Now()
	end := start.Add(duration)
	timer := time.NewTimer(duration)

	return &Countdown{sync.Mutex{}, start, end, timer, duration}
}

func (c *Countdown) Reset() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.end = time.Now().Add(c.duration)

	c.timer.Reset(c.duration)
}

func (c *Countdown) Stop() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.timer.Stop()
}

func (c *Countdown) Duration() time.Duration {
	return c.duration
}

func (c *Countdown) Remaining() time.Duration {
	c.mu.Lock()
	defer c.mu.Unlock()

	return time.Until(c.end)
}

func (c *Countdown) Elapsed() time.Duration {
	c.mu.Lock()
	defer c.mu.Unlock()

	return time.Since(c.start)
}

func (c *Countdown) Completed() <-chan time.Time {
	return c.timer.C
}
