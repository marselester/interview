package searching

import "sort"

// sum2020 finds the two entries in an array that sum to 2020
// and returns their product using brute force approach (n² running time).
// For example, in [1721 979 366 299 675 1456] entries 0 and 3 add up to 2020.
// The result should be 1721 * 299 = 514579.
// See https://adventofcode.com/2020/day/1.
func sum2020(a []int, target int) int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == target {
				return a[i] * a[j]
			}
		}
	}
	return -1
}

/*
sum2020matrix is like sum2020 but it looks up the elements in the matrix.

The idea is to sort the array and divide it in half.
The first half can represent rows, and another represent columns.
Together they form a matrix whose values are sums of rows/columns intersections.
For example, find elements that add up to 10 in [1 3 4 6 7 8] array.
The desired values are distributed close to diagonal.

       6  7  8
	1  7  8  9 <- start from here
	3  9 10 11
	4 10 11 12

Starting from a[0][2] compare the sum with the target value, e.g., 10.
Since 9 < 10, move down a[1][2].
Since 11 > 10, move to the left a[1][1] (found 10).
*/
func sum2020matrix(a []int, target int) int {
	sort.Ints(a)

	i, j := 0, len(a)-1
	var sum int
	for i < len(a) && j >= 0 {
		sum = a[i] + a[j]
		switch {
		// Found the elements that add up to the target.
		case sum == target:
			return a[i] * a[j]
		// Move down.
		case sum < target:
			i++
		// Move left.
		case sum > target:
			j--
		}
	}

	return -1
}

// sum2020binsearch is like sum2020 but it looks up the elements using binary search
// (n log n running time).
//
// Entries #3 (4) and #6 (8) add up to 12 (target) in the sorted array [1 2 3 4 6 7 8].
// The idea is to find 12-a[i] element using binary search, e.g., 12-4=8.
func sum2020binsearch(a []int, target int) int {
	sort.Ints(a)

	var j int
	for i := 0; i < len(a); i++ {
		j = BinarySearch(a, target-a[i])
		if j != -1 {
			return a[i] * a[j]
		}
	}

	return -1
}

// triplesum2020 finds the three entries in an array that sum to 2020
// and returns their product using brute force approach (cubic running time).
func triplesum2020(a []int, target int) int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			for k := j + 1; k < len(a); k++ {
				if a[i]+a[j]+a[k] == target {
					return a[i] * a[j] * a[k]
				}
			}
		}
	}
	return -1
}

// triplesum2020binsearch is like triplesum2020 but it has n² log n running time.
func triplesum2020binsearch(a []int, target int) int {
	sort.Ints(a)

	var k int
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			k = BinarySearch(a, target-a[i]-a[j])
			if k != -1 {
				return a[i] * a[j] * a[k]
			}
		}
	}

	return -1
}
