package linkedlist

import (
	"fmt"
)

func ExampleReverse() {
	ll := Node{
		Val: "A",
		Next: &Node{
			Val: "B",
			Next: &Node{
				Val: "C",
				Next: &Node{
					Val: "D",
				},
			},
		},
	}
	for n := Reverse(&ll); n != nil; n = n.Next {
		fmt.Println(n.Val)
	}
	// Output:
	// D
	// C
	// B
	// A
}
