package ratement

import "testing"

var testRates = map[int]int{
	1:   10,
	5:   60,
	10:  180,
	20:  480,
	50:  1440,
	100: 4320,
}

var rater = NewRater(testRates)

func TestRaterValue(t *testing.T) {
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

func TestRaterHas(t *testing.T) {
	tests := []struct {
		name     string
		id       int
		expected bool
	}{
		{"existing rate", 1, true},
		{"another existing rate", 5, true},
		{"non-existing rate", 3, false},
		{"another non-existing rate", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rater.Has(tt.id)

			if got != tt.expected {
				t.Errorf("Has(%d) = %v, want %v", tt.id, got, tt.expected)
			}
		})
	}
}

func TestRaterGet(t *testing.T) {
	tests := []struct {
		name     string
		id       int
		expected int
	}{
		{"existing rate", 10, 180},
		{"another existing rate", 20, 480},
		{"non-existing rate", 8, 0},
		{"another non-existing rate", 15, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rater.Get(tt.id)

			if got != tt.expected {
				t.Errorf("Get(%d) = %v, want %v", tt.id, got, tt.expected)
			}
		})
	}
}

func TestRaterSet(t *testing.T) {
	tests := []struct {
		name     string
		id       int
		expected int
	}{
		{"overwritten rate", 50, 1600},
		{"another overwritten rate", 100, 4800},
		{"newly added rate", 200, 10000},
		{"newly added rate", 500, 30000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rater.Set(tt.id, tt.expected)

			got := rater.Get(tt.id)

			if got != tt.expected {
				t.Errorf("Set(%d, %d) = %v, want %v", tt.id, tt.expected, got, tt.expected)
			}
		})
	}
}

func TestRaterDelete(t *testing.T) {
	tests := []struct {
		name     string
		id       int
		expected bool
	}{
		{"existing rate", 1, true},
		{"non-existing rate", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rater.Delete(tt.id)

			got := rater.Has(tt.id)

			if got == tt.expected {
				t.Errorf("Delete(%d) = %v, want %v", tt.id, got, tt.expected)
			}
		})
	}
}
