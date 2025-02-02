package eventuate

import (
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	started := time.Now()
	countdown := NewCountdown(2 * time.Second)

	<-countdown.Completed()

	if duration := countdown.Duration(); time.Since(started) < duration {
		t.Errorf("Expected time since started to be >= %v, got %v", duration, time.Since(started))
	}

	countdown.Reset()

	sleeper := 1 * time.Second

	time.Sleep(sleeper)

	if elapsed := countdown.Elapsed(); elapsed <= sleeper {
		t.Errorf("Expected elapsed time to be > %v, got %v", sleeper, elapsed)
	}

	countdown.Stop()

	if remaining := countdown.Remaining(); remaining <= 0 {
		t.Errorf("Expected remaining time to be > 0, got %v", remaining)
	}
}
