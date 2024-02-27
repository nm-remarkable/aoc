package advent

import (
	"testing"

	"github.com/stretchr/testify/require"
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
			require.Equal(t, canExist, tt.expect)
			require.Equal(t, tt.err, err != nil)
		})
	}
}
