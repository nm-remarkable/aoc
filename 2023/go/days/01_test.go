package advent

import (
	"testing"
)

func TestMatchDigit(t *testing.T) {
	str1 := "dvbjrbsixf4thre51oneightts7p"
	e1 := "6"
	if r1, err := MatchDigit(str1, false); err != nil || e1 != r1 {
		t.Fatalf(`First: Expected %s, got %s`, e1, r1)
	}

	e2 := "7"
	if r2, err := MatchDigit(str1, true); err != nil || e2 != r2 {
		t.Fatalf(`Last: Expected %s, got %s`, e2, r2)
	}

	str2 := "dvbbjrbfjsxffnjlhfdqthree51oneighttsp"

	e3 := "3"
	if r3, err := MatchDigit(str2, false); err != nil || e3 != r3 {
		t.Fatalf(`First: Expected %s, got %s`, e3, r3)
	}

	e4 := "8"
	if r4, err := MatchDigit(str2, true); err != nil || e4 != r4 {
		t.Fatalf(`Last: Expected %s, got %s`, e4, r4)
	}

	// Make sure the string we send is not modified
	if r5, err := MatchDigit(str2, false); err != nil || e3 != r5 {
		t.Fatalf(`First: Expected %s, got %s`, e3, r5)
	}
}
