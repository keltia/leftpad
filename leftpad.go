// leftpad.go

/*
 This package implements the leftpad function, inspired by the NPM (JS)
 package of the same name.

 Two functions are defined:

     import "leftpad"

     func LefPad(s, n)

     func LeftPadStr(s, n, c)

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

	prefix := strings.Repeat(c, toAdd)
	final := prefix + s
	return final, nil
}

func LeftPad(s string, n int) (string, error) {
    return doLeftPad(s, n, " ")
}


func LeftPadStr(s string, n int, c string) (string, error) {
	final, err := doLeftPad(s, n, c)
	return final, err
}
