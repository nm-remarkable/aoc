package advent

import (
	"fmt"
	"testing"
)

func TestCanHaveResult(t *testing.T) {
	bag := RGBag{10, 10, 10}

	fmt.Println("Bag has: ", bag)
	if isEnough, err := bag.gameIsValid("3 blue, 4 red"); !isEnough && err == nil {
		t.Fatalf("True: 3 blue, 4 red")
	}

	if isEnough, err := bag.gameIsValid("1 red, 2 green, 6 blue"); !isEnough && err == nil {
		t.Fatalf("True: 1 red, 2 green, 6 blue")
	}

	if isEnough, err := bag.gameIsValid("1 red, 11 green, 6 blue"); isEnough && err == nil {
		t.Fatalf("False: 1 red, 11 green, 6 blue")
	}

	if isEnough, err := bag.gameIsValid("100 red"); isEnough && err == nil {
		t.Fatalf("False: 100 red")
	}

	if _, err := bag.gameIsValid("dainodboasibda"); err == nil {
		t.Fatalf("String should not be parseable. Should throw error")
	}
}
