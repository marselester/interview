package array

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRotate90Counterclockwise(t *testing.T) {
	tt := map[string]struct {
		input [][]int
		want  [][]int
	}{
		"1x1": {
			input: [][]int{{1}},
			want:  [][]int{{1}},
		},
		"1x2": {
			input: [][]int{{1, 2}},
			want:  [][]int{{2}, {1}},
		},
		"3x2": {
			input: [][]int{{1, 2}, {3, 4}, {5, 6}},
			want:  [][]int{{2, 4, 6}, {1, 3, 5}},
		},
		"3x3": {
			input: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:  [][]int{{3, 6, 9}, {2, 5, 8}, {1, 4, 7}},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := Rotate90Counterclockwise(tc.input)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
