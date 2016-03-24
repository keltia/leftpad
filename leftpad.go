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
	"errors"
	"fmt"
	"strings"
)

var ErrInvalidChar = errors.New("Invalid character")

func doLeftPad(s string, n int, c string) (string, error) {
	if n < 0 {
		return "", errors.New(fmt.Sprintf("Invalid length %d", n))
	}

	if len(c) != 1 {
		return "", ErrInvalidChar
	}

	toAdd := n - len(s)
	if toAdd <= 0 {
		return s, nil
	}

	return strings.Repeat(c, toAdd) + s, nil
}

func LeftPad(s string, n int) (string, error) {
	return doLeftPad(s, n, " ")
}

func LeftPadStr(s string, n int, c string) (string, error) {
	return doLeftPad(s, n, c)
}
