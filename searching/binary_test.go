package searching

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tt := map[string]struct {
		input []int
		key   int
		want  int
	}{
		"leftmost": {
			input: []int{1, 3, 4, 6, 7, 8},
			key:   1,
			want:  0,
		},
		"rightmost": {
			input: []int{1, 3, 4, 6, 7, 8},
			key:   8,
			want:  5,
		},
		"not found": {
			input: []int{1, 3, 4, 6, 7, 8},
			key:   -10,
			want:  -1,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := BinarySearch(tc.input, tc.key)
			if tc.want != got {
				t.Errorf("expected %d got %d", tc.want, got)
			}
		})
	}
}
