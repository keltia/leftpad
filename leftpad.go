package leftpad

import (
	"fmt"
	"strings"
)

// Pad left-pads s with spaces, to length n.
// If n is smaller than s, Pad is a no-op.
func Pad(s string, n int) (string, error) {
	return PadChar(s, n, ' ')
}

// PadChar left-pads s with the rune r, to length n.
// If n is smaller than s, PadChar is a no-op.
func PadChar(s string, n int, r rune) (string, error) {
	if n < 0 {
		return "", fmt.Errorf("invalid length %d", n)
	}
	if len(s) > n {
		return s, nil
	}
	return strings.Repeat(string(r), n-len(s)) + s, nil
}
