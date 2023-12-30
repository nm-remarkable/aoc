package advent

import (
	"testing"
)

func TestMatchDigit(t *testing.T) {
	str := "dvbbjrbfjsxffnjlhfdqthree51oneighttsp"

	expectedFirst := "3"
	resultFirst := MatchDigit(str, false)
	if expectedFirst != resultFirst {
		t.Fatalf(`First: Expected %s, got %s`, expectedFirst, resultFirst)
	}

	expectedLast := "8"
	resultLast := MatchDigit(str, true)
	if expectedLast != resultLast {
		t.Fatalf(`Last: Expected %s, got %s`, expectedLast, resultLast)
	}

	str = "dvbjrbsixf4thre51oneightts7p"
	e1 := "6"
	r1 := MatchDigit(str, false)
	if e1 != r1 {
		t.Fatalf(`First: Expected %s, got %s`, e1, r1)
	}

	e2 := "7"
	r2 := MatchDigit(str, true)
	if e2 != r2 {
		t.Fatalf(`Last: Expected %s, got %s`, e2, r2)
	}
}
