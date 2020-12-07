package sorting

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMergesort(t *testing.T) {
	tt := map[string]struct {
		input []float64
		want  []float64
	}{
		"linearithmic-worst": {
			input: []float64{4, 9, 8, 3, 4, 7, 7, 0},
			want:  []float64{9, 8, 7, 7, 4, 4, 3, 0},
		},
		"linearithmic-duplicates": {
			input: []float64{0, 9, 0, 9, 0},
			want:  []float64{9, 9, 0, 0, 0},
		},
		"linear-ordered": {
			input: []float64{9, 8, 7, 7, 4, 4, 3, 0},
			want:  []float64{9, 8, 7, 7, 4, 4, 3, 0},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			Mergesort(tc.input)
			if diff := cmp.Diff(tc.want, tc.input); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
