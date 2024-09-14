package main

import (
	"testing"
)

func TestCanHaveResult(t *testing.T) {
	t.Parallel()

	tests := []struct {
		numbers  string
		expected []int
	}{
		{
			numbers:  "97 96 49 78 26 58 27 77 69  9 39 88 53 10  2 29 61 62 48 87 18 44 74 34 11",
			expected: []int{97, 96, 49, 78, 26, 58, 27, 77, 69, 9, 39, 88, 53, 10, 2, 29, 61, 62, 48, 87, 18, 44, 74, 34, 11},
		},
	}

	testFunc := convert
	for _, tt := range tests {
		tt := tt
		t.Run(tt.numbers, func(t *testing.T) {
			t.Parallel()
			nn, err := testFunc(tt.numbers)
			if err != nil {
				t.Errorf("Error: %s", err)
			}
			for i := range nn {
				if nn[i] != tt.expected[i] {
					t.Errorf("Expected: %d, got %d", tt.expected[i], nn[i])
				}
			}
		})
	}
}
