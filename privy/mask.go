package privy

import (
	"fmt"
	"strings"
	"unicode"
)

func Mask(s string, maskChar rune, visibleStart, visibleEnd int) string {
	// Use a slice of runes to handle multi-byte characters correctly.
	runes := []rune(s)
	length := len(runes)

	// If the string is shorter or equal to the total visible characters,
	// there's nothing to mask, so return the original string.
	if length <= visibleStart+visibleEnd {
		return s
	}

	// Create the masked middle part.
	maskLen := length - visibleStart - visibleEnd
	maskedPart := strings.Repeat(string(maskChar), maskLen)

	// Combine the start, masked, and end parts.
	return fmt.Sprintf("%s%s%s",
		string(runes[:visibleStart]),
		maskedPart,
		string(runes[length-visibleEnd:]),
	)
}
func MaskDigits(s string, maskChar rune, visibleStart, visibleEnd int) string {
	// First, count the total number of digits in the string.
	totalDigits := 0
	for _, r := range s {
		if unicode.IsDigit(r) {
			totalDigits++
		}
	}

	// If the number of digits is too small to be masked, return the original string.
	if totalDigits <= visibleStart+visibleEnd {
		return s
	}

	// Use a strings.Builder for efficient string construction.
	var builder strings.Builder
	builder.Grow(len(s)) // Pre-allocate capacity for performance.

	digitIndex := 0
	for _, r := range s {
		if unicode.IsDigit(r) {
			digitIndex++
			// Determine if the current digit should be masked.
			if digitIndex > visibleStart && digitIndex <= totalDigits-visibleEnd {
				builder.WriteRune(maskChar)
			} else {
				builder.WriteRune(r)
			}
		} else {
			// Not a digit, so append the character as is.
			builder.WriteRune(r)
		}
	}

	return builder.String()
}

func MaskEmail(s string, maskChar rune, visibleStart, visibleEnd int) string {
	atIndex := strings.LastIndex(s, "@")

	// If there's no '@' or it's the first character
	if atIndex <= 0 {
		// it's not a standard email. Apply a generic mask
		return Mask(s, maskChar, visibleStart, visibleEnd)
	}

	localPart := s[:atIndex]
	domain := s[atIndex:] // Includes the '@'
	localRunes := []rune(localPart)
	localLen := len(localRunes)
	visibleTotal := visibleStart + visibleEnd
	if localLen <= visibleTotal {
		// Handle the edge cases for very short local parts.
		switch localLen {
		case 1:
			// Entirely mask
			visibleStart = 0
			visibleEnd = 0
		case 2, 3:
			// Show first char only
			visibleStart = 1
			visibleEnd = 0
		case 4:
			// Show first and last chars
			visibleStart = 1
			visibleEnd = 1
		default:
			// Reduce visible chars by 1 on each side
			visibleStart -= 1
			visibleEnd -= 1
		}
	}

	return Mask(localPart, maskChar, visibleStart, visibleEnd) + domain
}
