package privy

import (
	"testing"
)

func TestMasking(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		maskChar     rune
		visibleStart int
		visibleEnd   int
		expected     string
	}{
		{
			name:         "Standard Case",
			input:        "abcdefghijkl",
			maskChar:     '*',
			visibleStart: 3,
			visibleEnd:   4,
			expected:     "abc*****ijkl",
		},
		{
			name:         "Too Short to Mask",
			input:        "abcdefg",
			maskChar:     '*',
			visibleStart: 3,
			visibleEnd:   4,
			expected:     "abcdefg",
			// length (7) == visibleStart + visibleEnd (7), so no mask
		},
		{
			name:         "Shorter Than Visible Parts",
			input:        "abc",
			maskChar:     '*',
			visibleStart: 3,
			visibleEnd:   4,
			expected:     "abc",
			// length (3) < visibleStart + visibleEnd (7), so no mask
		},
		{
			name:         "Empty String",
			input:        "",
			maskChar:     '*',
			visibleStart: 2,
			visibleEnd:   2,
			expected:     "",
		},
		{
			name:         "Mask Entire String",
			input:        "sensitive-data",
			maskChar:     '#',
			visibleStart: 0,
			visibleEnd:   0,
			expected:     "##############",
		},
		{
			name:         "Only Start Visible",
			input:        "abcdefghijkl",
			maskChar:     '*',
			visibleStart: 3,
			visibleEnd:   0,
			expected:     "abc*********",
		},
		{
			name:         "Only End Visible",
			input:        "abcdefghijkl",
			maskChar:     '*',
			visibleStart: 0,
			visibleEnd:   4,
			expected:     "********ijkl",
		},
		{
			name:         "Unicode String",
			input:        "你好世界12345",
			maskChar:     '•',
			visibleStart: 2,
			visibleEnd:   2,
			expected:     "你好•••••45",
			// "Hello World 12345"
			// Correctly handles multi-byte characters
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Mask(test.input, test.maskChar, test.visibleStart, test.visibleEnd)

			if got != test.expected {
				t.Errorf("Mask(%q) = %q, want %q", test.input, got, test.expected)
			}
		})
	}
}

func TestMaskDigits(t *testing.T) {
	testCases := []struct {
		name         string
		input        string
		maskChar     rune
		visibleStart int
		visibleEnd   int
		expected     string
	}{
		{
			name:         "Standard Case",
			input:        "8885551212",
			maskChar:     '#',
			visibleStart: 3,
			visibleEnd:   4,
			expected:     "888###1212",
		},
		{
			name:         "Too Short to Mask",
			input:        "1234567",
			maskChar:     '#',
			visibleStart: 3,
			visibleEnd:   4,
			expected:     "1234567",
		},
		{
			name:         "Shorter Than Visible Parts",
			input:        "123",
			maskChar:     '#',
			visibleStart: 3,
			visibleEnd:   4,
			expected:     "123",
		},
		{
			name:         "Empty String",
			input:        "",
			maskChar:     '#',
			visibleStart: 2,
			visibleEnd:   2,
			expected:     "",
		},
		{
			name:         "Formatted with Dashes",
			input:        "888-555-1212",
			maskChar:     '*',
			visibleStart: 4,
			visibleEnd:   3,
			expected:     "888-5**-*212",
		},
		{
			name:         "Formatted with Dots",
			input:        "888.555.1212",
			maskChar:     '*',
			visibleStart: 3,
			visibleEnd:   4,
			expected:     "888.***.1212",
		},
		{
			name:         "Phone with Country Code",
			input:        "+15551234567",
			maskChar:     'x',
			visibleStart: 4,
			visibleEnd:   4,
			expected:     "+1555xxx4567",
		},
		{
			name:         "With Parens and Dash",
			input:        "+1 (555) 123-4567",
			maskChar:     'x',
			visibleStart: 4,
			visibleEnd:   3,
			expected:     "+1 (555) xxx-x567",
		},
		{
			name:         "Country Code with Spaces",
			input:        "+49 176 12345678",
			maskChar:     'x',
			visibleStart: 4,
			visibleEnd:   4,
			expected:     "+49 17x xxxx5678",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := MaskDigits(tc.input, tc.maskChar, tc.visibleStart, tc.visibleEnd)

			if got != tc.expected {
				t.Errorf("MaskDigits() = %q, want %q", got, tc.expected)
			}
		})
	}
}

func TestMaskEmail(t *testing.T) {
	testCases := []struct {
		name         string
		input        string
		maskChar     rune
		visibleStart int
		visibleEnd   int
		expected     string
	}{
		{
			name:         "Standard Case",
			input:        "john.doe@example.com",
			maskChar:     '*',
			visibleStart: 3,
			visibleEnd:   2,
			expected:     "joh***oe@example.com",
		},
		{
			name:         "Long Local Part",
			input:        "jane_roe_here@gmail.com",
			maskChar:     '*',
			visibleStart: 4,
			visibleEnd:   3,
			expected:     "jane******ere@gmail.com",
		},
		{
			name:         "Shorter Than Visible Parts",
			input:        "testing@gmail.com",
			maskChar:     '*',
			visibleStart: 4,
			visibleEnd:   3,
			expected:     "tes**ng@gmail.com",
			// one less visible char each side
		},
		{
			name:         "Short Local Part (4 chars)",
			input:        "juan@web.dev",
			maskChar:     '#',
			visibleStart: 3,
			visibleEnd:   2,
			expected:     "j##n@web.dev",
		},
		{
			name:         "Short Local Part (3 chars)",
			input:        "bob@web.dev",
			maskChar:     '*',
			visibleStart: 3,
			visibleEnd:   2,
			expected:     "b**@web.dev",
		},
		{
			name:         "Short Local Part (2 chars)",
			input:        "me@web.dev",
			maskChar:     '*',
			visibleStart: 3,
			visibleEnd:   2,
			expected:     "m*@web.dev",
		},
		{
			name:         "Short Local Part (1 char)",
			input:        "a@b.co",
			maskChar:     '*',
			visibleStart: 3,
			visibleEnd:   2,
			expected:     "*@b.co",
		},
		{
			name:         "With Subdomain",
			input:        "support@help.my-company.co.uk",
			maskChar:     '*',
			visibleStart: 2,
			visibleEnd:   1,
			expected:     "su****t@help.my-company.co.uk",
		},
		{
			name:         "Without @ (Fallback Case)",
			input:        "not-an-email",
			maskChar:     '*',
			visibleStart: 3,
			visibleEnd:   2,
			expected:     "not*******il",
			// Falls back to default masking
		},
		{
			name:         "Starts @ (Fallback Case)",
			input:        "@not-an-email",
			maskChar:     '*',
			visibleStart: 3,
			visibleEnd:   2,
			expected:     "@no********il",
			// Falls back to default masking
		},
		{
			name:         "Empty String",
			input:        "",
			maskChar:     '*',
			visibleStart: 3,
			visibleEnd:   2,
			expected:     "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := MaskEmail(tc.input, tc.maskChar, tc.visibleStart, tc.visibleEnd)

			if got != tc.expected {
				t.Errorf("MaskEmail() = %q, want %q", got, tc.expected)
			}
		})
	}
}
