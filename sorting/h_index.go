package sorting

import (
	"sort"
)

// Hindex calculates h-index metric that measures both productivity and citation impact of a researcher.
// It's the largest number that each paper was cited at least h times.
// For example, papers A, B, C, D, E, F, G, H, I were cited 1, 4, 1, 4, 2, 1, 3, 5, 6 times respectively.
// The h-index is 4 (it corresponds to papers B, D, H, I).
//
// Iterate throught sorted citations (1, 1, 1, 2, 3, 4, 4, 5, 6)
// and stop at the first 4 (i=5) since 4 >= number of remaining papers.
// Running time complexity is O(n * log n) where sorting is log n and a for loop is n.
func Hindex(citations []int) int {
	sort.Ints(citations)
	for i, c := range citations {
		if c >= len(citations)-i {
			return len(citations) - i
		}
	}
	return -1
}
