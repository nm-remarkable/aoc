package advent

import (
	"reflect"
	"testing"
)

func TestCanHaveResult(t *testing.T) {
	t.Parallel()

	bag := RGBag{10, 10, 10}
	tests := []struct {
		name   string
		given  string
		expect bool
		err    bool
	}{
		{
			name:   "Result is valid 2 different colors",
			given:  "3 blue, 4 red",
			expect: true,
			err:    false,
		},
		{
			name:   "Result is valid 3 different colors",
			given:  "1 red, 2 green, 6 blue",
			expect: true,
			err:    false,
		},
		{
			name:   "Too many balls with multiple colors",
			given:  "1 red, 11 green, 6 blue",
			expect: false,
			err:    false,
		},
		{
			name:   "Too many balls with only one color",
			given:  "100 red",
			expect: false,
			err:    false,
		},
		{
			name:   "Bad input, only alphabetical",
			given:  "fnaoisndiasnda",
			expect: false,
			err:    true,
		},
		{
			name:   "Bad input, mix but no colors",
			given:  "100 apples",
			expect: false,
			err:    true,
		},
	}

	testFunc := bag.gameIsValid
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			canExist, err := testFunc(tt.given)
			if tt.expect != canExist {
				t.Errorf("Expected: %t, got %t", tt.expect, canExist)
			}
			if tt.err != (err != nil) {
				t.Errorf("Error: %s", err)
			}
		})
	}
}

func TestMinimum(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		given  string
		expect RGBag
		err    bool
	}{
		{
			given:  "3 blue, 4 red",
			expect: RGBag{4, 0, 3},
			err:    false,
		},
		{
			given:  "1 red, 2 green, 6 blue",
			expect: RGBag{1, 2, 6},
			err:    false,
		},
		{
			given:  "1 red, 11 green, 6 blue",
			expect: RGBag{1, 11, 6},
			err:    false,
		},
		{
			given:  "1 red; 1 green; 1 blue",
			expect: RGBag{1, 1, 1},
			err:    false,
		},
		{
			given:  "1 red; 2 red; 3 red; 1 blue",
			expect: RGBag{3, 0, 1},
			err:    false,
		},
		{
			given:  "1 red, 2 green; 2 red, 1 green",
			expect: RGBag{2, 2, 0},
			err:    false,
		},
		{
			given:  "100 red",
			expect: RGBag{100, 0, 0},
			err:    false,
		},
		{
			given:  "fnaoisndiasnda",
			expect: RGBag{},
			err:    true,
		},
		{
			given:  "100 apples",
			expect: RGBag{},
			err:    true,
		},
	}

	testFunc := minimum
	for _, tt := range tests {
		tt := tt
		t.Run(tt.given, func(t *testing.T) {
			t.Parallel()
			minimum, err := testFunc(tt.given)
			if !reflect.DeepEqual(minimum, tt.expect) {
				t.Errorf("unexpected value: %+v, expected: %+v", minimum, tt.expect)
			}
			if tt.err != (err != nil) {
				t.Errorf("Error: %s", err)
			}
		})
	}
}
