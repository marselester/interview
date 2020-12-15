package linkedlist

// Node represents a node in a linked list.
type Node struct {
	Val  string
	Next *Node
}

/*
Reverse reverses linked list starting from cur node in linear time.
The idea is to point the current node c to the previous node p, i.e., A->nil, B->A, C->B, D->C.

	  A B C D
	p c n
	  p c n
	    p c n
	      p c n
	        p c n
*/
func Reverse(cur *Node) *Node {
	var (
		prev *Node
		next *Node
	)
	// A B C
	// p c n
	for cur != nil {
		// Store next node C which will become current in the following iteration.
		next = cur.Next
		// Current node B points to the previous node A.
		cur.Next = prev
		// Advance previous and current node pointers for the following iteration.
		prev = cur
		cur = next
	}
	return prev
}
