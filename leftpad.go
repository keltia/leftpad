/*
 This package implements the leftpad function, inspired by the NPM (JS)
 package of the same name.

 Two functions are defined:

 import "leftpad"

 // pad with spaces
 str, err := LeftPad(s, n)

 // pad with specified character
 str, err := func LeftPadStr(s, n, c)

*/
package leftpad

import (
	"fmt"
	"strings"
)

// LeftPad left-pads s with spaces, to length n.
// If n is smaller than s, LeftPad is a no-op.
func LeftPad(s string, n int) (string, error) {
	return LeftPadStr(s, n, ' ')
}

// LeftPadStr left-pads s with the rune r, to length n.
// If n is smaller than s, LeftPadStr is a no-op.
func LeftPadStr(s string, n int, r rune) (string, error) {
	if n < 0 {
		return "", fmt.Errorf("invalid length %d", n)
	}
	if len(s) > n {
		return s, nil
	}
	return strings.Repeat(string(r), n-len(s)) + s, nil
}
