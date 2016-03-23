// leftpad_test.go

/*
  These are the tests for the leftpad package.
 */
package leftpad

import (
	"testing"
	"fmt"
)

var (
	basic = "foo"
)

func TestLeftPad(t *testing.T) {
	res, err := LeftPad(basic, 2)
	if res != basic || err != nil {
		t.Errorf(fmt.Sprintf("Strings should match! res='%s'", res))
	}

	res, err = LeftPad(basic, 3)
	if res != basic || err != nil {
		t.Errorf(fmt.Sprintf("Strings should match! res='%s'", res))
	}

	res, err = LeftPad(basic, 5)
	if res != "  foo" || err != nil {
		t.Errorf(fmt.Sprintf("Error: '%s' should be '  foo'", res))
	}
}

func TestLeftPadStr(t *testing.T) {
	res, err := LeftPadStr(basic, 2, "X")
	if res != basic || err != nil {
		t.Errorf(fmt.Sprintf("Strings should match! res='%s'", res))
	}

	res, err = LeftPadStr(basic, 3, "X")
	if res != basic || err != nil {
		t.Errorf(fmt.Sprintf("Strings should match! res='%s'", res))
	}

	res, err = LeftPadStr(basic, 5, "X")
	if res != "XXfoo" || err != nil {
		t.Errorf(fmt.Sprintf("Error: '%s' should be 'XXfoo'", res))
	}
}
