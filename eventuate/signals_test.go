package eventuate

import (
	"testing"
	"time"
)

func TestSignals(t *testing.T) {
	signals := NewSignals(10 * time.Millisecond)
	counter := make(map[int]int)

	go func() {
		final := false

		for {
			select {
			case count := <-signals.Every(1):
				counter[1] = count
			case count := <-signals.Every(2):
				counter[2] = count
			case count := <-signals.Every(3):
				counter[3] = count
			case count := <-signals.Every(5):
				counter[5] = count
			case <-signals.Every(10):
				if final {
					signals.Stop()
				}

				signals.Reset()
				final = true
			}
		}
	}()

	time.Sleep(200 * time.Millisecond)

	tests := []struct {
		name     string
		nth      int
		expected int
	}{
		// expect only half due to reset
		{"Every 1*10ms for 200ms", 1, 10},
		{"Every 2*10ms for 200ms", 2, 5},
		{"Every 3*10ms for 200ms", 3, 3},
		{"Every 5*10ms for 200ms", 5, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := counter[tt.nth]; got != tt.expected {
				t.Errorf("Signals.Every(%d) = %v, want %v", tt.expected, got, tt.expected)
			}
		})
	}
}
