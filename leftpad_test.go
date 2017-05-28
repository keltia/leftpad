package leftpad

import "testing"

func TestPad(t *testing.T) {
	for _, testcase := range []struct {
		s    string
		n    int
		want string
	}{
		{"foo", 2, "foo"},
		{"foo", 3, "foo"},
		{"foo", 4, " foo"},
		{"foo", 5, "  foo"},
	} {
		have, err := Pad(testcase.s, testcase.n)
		if err != nil {
			t.Errorf("Pad(%q, %d): %v", testcase.s, testcase.n, err)
			continue
		}
		if want := testcase.want; want != have {
			t.Errorf("Pad(%q, %d): want %q, have %q", testcase.s, testcase.n, want, have)
			continue
		}
	}
}

func TestPadChar(t *testing.T) {
	for _, testcase := range []struct {
		s    string
		n    int
		r    rune
		want string
	}{
		{"foo", 2, 'X', "foo"},
		{"foo", 3, 'X', "foo"},
		{"foo", 4, 'X', "Xfoo"},
		{"foo", 5, 'X', "XXfoo"},
	} {
		have, err := PadChar(testcase.s, testcase.n, testcase.r)
		if err != nil {
			t.Errorf("PadChar(%q, %d, %c): %v", testcase.s, testcase.n, testcase.r, err)
			continue
		}
		if want := testcase.want; want != have {
			t.Errorf("Pad(%q, %d, %c): want %q, have %q", testcase.s, testcase.n, testcase.r, want, have)
			continue
		}
	}
	have, err := PadChar("foo", -1, 'X')
	if err == nil || have != ""{
		t.Error("PadChar with negative number should be an error")
	}
}
