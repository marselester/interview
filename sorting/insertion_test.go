package sorting

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInsertionsort(t *testing.T) {
	tt := map[string]struct {
		input []float64
		want  []float64
	}{
		"n²-random": {
			input: []float64{9, 4, 8, 3, 7},
			want:  []float64{9, 8, 7, 4, 3},
		},
		"n²-reverse-order": {
			input: []float64{7, 3, 8, 4, 9},
			want:  []float64{9, 8, 7, 4, 3},
		},
		"linear-sorted": {
			input: []float64{9, 8, 7, 4, 3},
			want:  []float64{9, 8, 7, 4, 3},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			Insertionsort(tc.input)
			if diff := cmp.Diff(tc.want, tc.input); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
