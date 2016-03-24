// leftpad_test.go

/*
  These are the tests for the leftpad package.
*/
package leftpad

import (
	"fmt"
	"testing"
)

var (
	basic = "foo"
)

func TestPad(t *testing.T) {
	res, err := Pad(basic, 2)
	if res != basic || err != nil {
		t.Errorf(fmt.Sprintf("Strings should match! res='%s'", res))
	}

	res, err = Pad(basic, 3)
	if res != basic || err != nil {
		t.Errorf(fmt.Sprintf("Strings should match! res='%s'", res))
	}

	res, err = Pad(basic, 5)
	if res != "  foo" || err != nil {
		t.Errorf(fmt.Sprintf("Error: '%s' should be '  foo'", res))
	}
}

func TestPadChar(t *testing.T) {
	res, err := PadChar(basic, 2, 'X')
	if res != basic || err != nil {
		t.Errorf(fmt.Sprintf("Strings should match! res='%s'", res))
	}

	res, err = PadChar(basic, 3, 'X')
	if res != basic || err != nil {
		t.Errorf(fmt.Sprintf("Strings should match! res='%s'", res))
	}

	res, err = PadChar(basic, 5, 'X')
	if res != "XXfoo" || err != nil {
		t.Errorf(fmt.Sprintf("Error: '%s' should be 'XXfoo'", res))
	}
}
