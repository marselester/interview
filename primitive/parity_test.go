package primitive

import "testing"

func TestParityNaive(t *testing.T) {
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
			want:  0,
		},
		"0000 1001": {
			input: 9,
			want:  0,
		},
		"0001 0111 1101 0110 1110 0011": {
			input: 1562339,
			want:  0,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := ParityNaive(tc.input)
			if tc.want != got {
				t.Errorf("expected %d got %d", tc.want, got)
			}
		})
	}
}

func TestParityNaiveToggle(t *testing.T) {
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
			want:  0,
		},
		"0000 1001": {
			input: 9,
			want:  0,
		},
		"0001 0111 1101 0110 1110 0011": {
			input: 1562339,
			want:  0,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := ParityNaiveToggle(tc.input)
			if tc.want != got {
				t.Errorf("expected %d got %d", tc.want, got)
			}
		})
	}
}

func TestParityKernighan(t *testing.T) {
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
			want:  0,
		},
		"0000 1001": {
			input: 9,
			want:  0,
		},
		"0001 0111 1101 0110 1110 0011": {
			input: 1562339,
			want:  0,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := ParityKernighan(tc.input)
			if tc.want != got {
				t.Errorf("expected %d got %d", tc.want, got)
			}
		})
	}
}

func TestParityXor(t *testing.T) {
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
			want:  0,
		},
		"0000 1001": {
			input: 9,
			want:  0,
		},
		"0001 0111 1101 0110 1110 0011": {
			input: 1562339,
			want:  0,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := ParityXor(tc.input)
			if tc.want != got {
				t.Errorf("expected %d got %d", tc.want, got)
			}
		})
	}
}

func TestParityLookupTable(t *testing.T) {
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
			want:  0,
		},
		"0000 1001": {
			input: 9,
			want:  0,
		},
		"0001 0111 1101 0110 1110 0011": {
			input: 1562339,
			want:  0,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := ParityLookupTable(tc.input)
			if tc.want != got {
				t.Errorf("expected %d got %d", tc.want, got)
			}
		})
	}
}

func BenchmarkParityNaive(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ParityNaive(1562339)
	}
}

func BenchmarkParityNaiveToggle(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ParityNaiveToggle(1562339)
	}
}

func BenchmarkParityKernighan(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ParityKernighan(1562339)
	}
}

func BenchmarkParityXor(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ParityXor(1562339)
	}
}

func BenchmarkParityLookupTable(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ParityLookupTable(1562339)
	}
}
