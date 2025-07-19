package privy

import (
	"testing"
)

func TestDecodeCF(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "Valid Email",
			// "test@example.com" XORed with key 0x55
			input:    "552130262115302d34382539307b363a38",
			expected: "test@example.com",
		},
		{
			name: "Another Valid Email",
			// "info@company.net" XORed with key 0xab
			input:    "abc2c5cdc4ebc8c4c6dbcac5d285c5cedf",
			expected: "info@company.net",
		},
		{
			name:     "Invalid Hex String",
			input:    "not-a-hex-string",
			expected: "",
		},
		{
			name:     "Empty String",
			input:    "",
			expected: "",
		},
		{
			name:     "String too short (only a key)",
			input:    "aa",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DecodeCF(tt.input)

			if got != tt.expected {
				t.Errorf("DecodeCF(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}
