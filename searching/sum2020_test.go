package searching

import (
	"testing"
)

var adventInput = []int{1768, 1847, 1905, 1713, 1826, 1846, 1824, 1976, 1687, 1867, 1665, 1606, 1946, 1886, 1858, 346, 1739, 1752, 1700, 1922, 1865, 1609, 1617, 1932, 1346, 1213, 1933, 834, 1598, 1191, 1979, 1756, 1216, 1820, 1792, 1537, 1341, 1390, 1709, 1458, 1808, 1885, 1679, 1977, 1869, 1614, 1938, 1622, 1868, 1844, 1969, 1822, 1510, 1994, 1337, 1883, 1519, 1766, 1554, 1825, 1828, 1972, 1380, 1878, 1345, 1469, 1794, 1898, 1805, 1911, 1913, 1910, 1318, 1862, 1921, 1753, 1823, 1896, 1316, 1381, 1430, 1962, 1958, 1702, 1923, 1993, 1789, 2002, 1788, 1970, 1955, 1887, 1870, 225, 1696, 1975, 699, 294, 1605, 1500, 1777, 1750, 1857, 1540, 1329, 1974, 1947, 1516, 1925, 1945, 350, 1669, 1775, 1536, 1871, 1917, 1249, 1971, 2009, 1585, 1986, 1701, 1832, 1754, 1195, 1697, 1941, 1919, 2006, 1667, 1816, 1765, 1631, 2003, 1861, 1000, 1791, 1786, 1843, 1939, 1951, 269, 1790, 1895, 1355, 1833, 1466, 1998, 1806, 1881, 1234, 1856, 1619, 1727, 1874, 1877, 195, 1783, 1797, 2010, 1764, 1863, 1852, 1841, 1892, 1562, 1650, 1942, 1695, 1730, 1965, 1632, 1981, 1900, 1991, 1884, 1278, 1062, 1394, 1999, 2000, 1827, 1873, 1926, 1434, 1802, 1579, 1879, 1671, 1549, 1875, 1838, 1338, 1864, 1718, 1800, 1928, 1749, 1990, 1705}

func TestSum2020(t *testing.T) {
	tt := map[string]struct {
		input  []int
		target int
		want   int
	}{
		"[1][4]": {
			input:  []int{1, 3, 4, 6, 7, 8},
			target: 10,
			want:   21,
		},
		"[0][5] advent": {
			input:  []int{1721, 979, 366, 299, 675, 1456},
			target: 2020,
			want:   514579,
		},
		"[3][4]": {
			input:  []int{1, 3, 4, 6, 7, 8},
			target: 13,
			want:   42,
		},
		"[0][1]": {
			input:  []int{1, 3, 4, 6, 7, 8},
			target: 4,
			want:   3,
		},
		"not found": {
			input:  []int{1, 3, 4, 6, 7, 8},
			target: 100,
			want:   -1,
		},
		"big": {
			input:  adventInput,
			target: 2020,
			want:   355875,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := sum2020(tc.input, tc.target)
			if tc.want != got {
				t.Errorf("expected %d got %d", tc.want, got)
			}
		})
	}
}

func TestSum2020matrix(t *testing.T) {
	tt := map[string]struct {
		input  []int
		target int
		want   int
	}{
		"[1][4]": {
			input:  []int{1, 3, 4, 6, 7, 8},
			target: 10,
			want:   21,
		},
		"[0][5] advent": {
			input:  []int{1721, 979, 366, 299, 675, 1456},
			target: 2020,
			want:   514579,
		},
		"[3][4]": {
			input:  []int{1, 3, 4, 6, 7, 8},
			target: 13,
			want:   42,
		},
		"[0][1]": {
			input:  []int{1, 3, 4, 6, 7, 8},
			target: 4,
			want:   3,
		},
		"not found": {
			input:  []int{1, 3, 4, 6, 7, 8},
			target: 100,
			want:   -1,
		},
		"big": {
			input:  adventInput,
			target: 2020,
			want:   355875,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := sum2020matrix(tc.input, tc.target)
			if tc.want != got {
				t.Errorf("expected %d got %d", tc.want, got)
			}
		})
	}
}

func TestSum2020binsearch(t *testing.T) {
	tt := map[string]struct {
		input  []int
		target int
		want   int
	}{
		"[1][4]": {
			input:  []int{1, 3, 4, 6, 7, 8},
			target: 10,
			want:   21,
		},
		"[0][5] advent": {
			input:  []int{1721, 979, 366, 299, 675, 1456},
			target: 2020,
			want:   514579,
		},
		"[3][4]": {
			input:  []int{1, 3, 4, 6, 7, 8},
			target: 13,
			want:   42,
		},
		"[0][1]": {
			input:  []int{1, 3, 4, 6, 7, 8},
			target: 4,
			want:   3,
		},
		"not found": {
			input:  []int{1, 3, 4, 6, 7, 8},
			target: 100,
			want:   -1,
		},
		"big": {
			input:  adventInput,
			target: 2020,
			want:   355875,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := sum2020binsearch(tc.input, tc.target)
			if tc.want != got {
				t.Errorf("expected %d got %d", tc.want, got)
			}
		})
	}
}

func BenchmarkSum2020(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sum2020(adventInput, 2020)
	}
}

func BenchmarkSum2020matrix(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sum2020matrix(adventInput, 2020)
	}
}

func BenchmarkSum2020binsearch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sum2020binsearch(adventInput, 2020)
	}
}

func TestTriplesum2020(t *testing.T) {
	tt := map[string]struct {
		input  []int
		target int
		want   int
	}{
		"[1][4][5]": {
			input:  []int{1, 3, 4, 6, 7, 8},
			target: 18,
			want:   168,
		},
		"[1][2][4] advent": {
			input:  []int{1721, 979, 366, 299, 675, 1456},
			target: 2020,
			want:   241861950,
		},
		"not found": {
			input:  []int{1, 3, 4, 6, 7, 8},
			target: 100,
			want:   -1,
		},
		"big": {
			input:  adventInput,
			target: 2020,
			want:   140379120,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := triplesum2020(tc.input, tc.target)
			if tc.want != got {
				t.Errorf("expected %d got %d", tc.want, got)
			}
		})
	}
}

func TestTriplesum2020binsearch(t *testing.T) {
	tt := map[string]struct {
		input  []int
		target int
		want   int
	}{
		"[1][4][5]": {
			input:  []int{1, 3, 4, 6, 7, 8},
			target: 18,
			want:   168,
		},
		"[1][2][4] advent": {
			input:  []int{1721, 979, 366, 299, 675, 1456},
			target: 2020,
			want:   241861950,
		},
		"not found": {
			input:  []int{1, 3, 4, 6, 7, 8},
			target: 100,
			want:   -1,
		},
		"big": {
			input:  adventInput,
			target: 2020,
			want:   140379120,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := triplesum2020binsearch(tc.input, tc.target)
			if tc.want != got {
				t.Errorf("expected %d got %d", tc.want, got)
			}
		})
	}
}

func BenchmarkTriplesumsum2020(b *testing.B) {
	for n := 0; n < b.N; n++ {
		triplesum2020(adventInput, 2020)
	}
}

func BenchmarkTriplesum2020binsearch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		triplesum2020binsearch(adventInput, 2020)
	}
}
