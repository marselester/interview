package array

import "testing"

func TestDutchFlag(t *testing.T) {
	tt := map[string]struct {
		input []int
		index int
		want  []int
	}{
		"0 1 2 _0_ 2 1 1": {
			input: []int{0, 1, 2, 0, 2, 1, 1},
			index: 3,
			want:  []int{0, 0, 1, 2, 2, 1, 1},
		},
		"0 1 _2_ 0 2 1 1": {
			input: []int{0, 1, 2, 0, 2, 1, 1},
			index: 2,
			want:  []int{0, 1, 0, 1, 1, 2, 2},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			DutchFlag(tc.input, tc.index)
			if !equal(tc.want, tc.input) {
				t.Errorf("expected %d got %d", tc.want, tc.input)
			}
		})
	}
}

func TestDutchFlagOnePass(t *testing.T) {
	tt := map[string]struct {
		input []int
		index int
		want  []int
	}{
		"0 1 2 _0_ 2 1 1": {
			input: []int{0, 1, 2, 0, 2, 1, 1},
			index: 3,
			want:  []int{0, 0, 2, 2, 1, 1, 1},
		},
		"0 1 _2_ 0 2 1 1": {
			input: []int{0, 1, 2, 0, 2, 1, 1},
			index: 2,
			want:  []int{0, 1, 0, 1, 1, 2, 2},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			DutchFlagOnePass(tc.input, tc.index)
			if !equal(tc.want, tc.input) {
				t.Errorf("expected %d got %d", tc.want, tc.input)
			}
		})
	}
}
