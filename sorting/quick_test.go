package sorting

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestQuicksort(t *testing.T) {
	tt := map[string]struct {
		input []float64
		want  []float64
	}{
		"small": {
			input: []float64{9, 4, 8, 3, 7},
			want:  []float64{9, 8, 7, 4, 3},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			Quicksort(tc.input)
			if diff := cmp.Diff(tc.want, tc.input); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
