package advent

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMatchDigit(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		input   string
		reverse bool
		expect  string
	}{
		{
			name:    "Left match numbered word",
			input:   "dvbjrbsixf4thre51oneightts7p",
			reverse: false,
			expect:  "6",
		},
		{
			name:    "Right match digit",
			input:   "dvbjrbsixf4thre51oneightts7p",
			reverse: true,
			expect:  "7",
		},
		{
			name:    "Left match digit",
			input:   "dvbbjrbfjsxffnjlhfdq2three51oneighttsp",
			reverse: false,
			expect:  "2",
		},
		{
			name:    "Left match (nineight)",
			input:   "dvbbjrbfjsnineighthree51oneighttsp",
			reverse: false,
			expect:  "9",
		},
		{
			name:    "Right match (oneight = 8)",
			input:   "dvbbjrbfjsxffnjlhfdqthree51oneighttsp",
			reverse: true,
			expect:  "8",
		},
	}
	testFunc := MatchDigit
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			n, err := testFunc(tt.input, tt.reverse)
			require.NoError(t, err)
			require.Equal(t, tt.expect, n, fmt.Sprintf(`Expected %s, got %s`, tt.input, n))
		})
	}
}
