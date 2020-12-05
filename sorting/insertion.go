package sorting

/*
Insertionsort sorts a in descending order using insertion sort algorithm (stable, in place).
The idea is to examine elements from i=1 till the end and
go backwards on each iteration to swap neighbors if sorting condition has changed.

	1 4 3 -1
	  i->
	<-j

	1) a[1] > a[0], swap(a[1], a[0]) => 4 1 3 -1
	2) a[2] > a[1], swap(a[2], a[1]) => 4 3 1 -1
	   a[1] < a[0]
	3) a[3] < a[2]

Running time is between n and n² (depends on order of items).
Insertion sort is an excellent method for partially sorted arrays and
is also a fine method for tiny arrays (~15 items).
These properties can be leveraged in intermediate stages of mergesort and quicksort.
*/
func Insertionsort(a []float64) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j] > a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			} else {
				break
			}
		}
	}
}
