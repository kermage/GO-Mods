package ratement

import "testing"

func TestRater(t *testing.T) {
	rater := NewRater(map[int]int{
		1:   10,
		5:   60,
		10:  180,
		20:  480,
		50:  1440,
		100: 4320,
	})
	tests := []struct {
		name     string
		amount   int
		expected int
	}{
		{"2 coins for 20 minutes", 2, 20},
		{"8 coins for 90 minutes", 8, 90},
		{"15 coins for 240 minutes", 15, 240},
		{"30 coins for 660 minutes", 30, 660},
		{"90 coins for 2400 minutes", 90, 2400},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rater.Value(tt.amount)

			if got != tt.expected {
				t.Errorf("Value(%d) = %v, want %v", tt.amount, got, tt.expected)
			}
		})
	}
}
