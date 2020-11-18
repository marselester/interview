package array

import "testing"

func TestEvenOdd(t *testing.T) {
	tt := map[string]struct {
		input []int
		want  []int
	}{
		"nil": {
			input: nil,
			want:  nil,
		},
		"one item": {
			input: []int{1},
			want:  []int{1},
		},
		"odd even": {
			input: []int{1, 2},
			want:  []int{2, 1},
		},
		"odd odd even": {
			input: []int{1, 3, 6},
			want:  []int{6, 3, 1},
		},
		"1 3 2 5 6": {
			input: []int{1, 3, 2, 5, 6},
			want:  []int{6, 2, 5, 3, 1},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			EvenOdd(tc.input)
			if !equal(tc.want, tc.input) {
				t.Errorf("expected %d got %d", tc.want, tc.input)
			}
		})
	}
}

func equal(a1, a2 []int) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}
