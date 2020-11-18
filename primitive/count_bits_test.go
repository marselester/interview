package primitive

import "testing"

func TestCountBitsNaive(t *testing.T) {
	tt := map[string]struct {
		input int
		want  int
	}{
		"0000 0001": {
			input: 1,
			want:  1,
		},
		"0000 0011": {
			input: 3,
			want:  2,
		},
		"0000 1001": {
			input: 9,
			want:  2,
		},
		"0001 0111 1101 0110 1110 0011": {
			input: 1562339,
			want:  14,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := CountBitsNaive(tc.input)
			if tc.want != got {
				t.Errorf("expected %d got %d", tc.want, got)
			}
		})
	}
}

func TestCountBitsKernighan(t *testing.T) {
	tt := map[string]struct {
		input int
		want  int
	}{
		"0000 0001": {
			input: 1,
			want:  1,
		},
		"0000 0011": {
			input: 3,
			want:  2,
		},
		"0000 1001": {
			input: 9,
			want:  2,
		},
		"0001 0111 1101 0110 1110 0011": {
			input: 1562339,
			want:  14,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := CountBitsKernighan(tc.input)
			if tc.want != got {
				t.Errorf("expected %d got %d", tc.want, got)
			}
		})
	}
}

func BenchmarkCountBitsNaive(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CountBitsNaive(1562339)
	}
}

func BenchmarkCountBitsKernighan(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CountBitsKernighan(1562339)
	}
}
