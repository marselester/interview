package sorting

import "testing"

func TestHindex(t *testing.T) {
	tt := map[string]struct {
		input []int
		want  int
	}{
		"EPI book example": {
			input: []int{1, 1, 1, 2, 3, 4, 4, 5, 6},
			want:  4,
		},
		"blank": {
			input: []int{},
			want:  -1,
		},
		"nil": {
			input: nil,
			want:  -1,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := Hindex(tc.input)
			if tc.want != got {
				t.Errorf("expected %d got %d", tc.want, got)
			}
		})
	}
}
