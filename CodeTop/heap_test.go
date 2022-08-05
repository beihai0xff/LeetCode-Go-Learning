package CodeTop

import (
	"container/heap"
	"fmt"
)

// This example inserts several ints into an MaxHeap, checks the maximum,
// and removes them in order of priority.
func Example_intHeap() {
	h := &MaxHeap{[]int{1, 3, 2}}
	heap.Init(h)
	heap.Push(h, 7)
	fmt.Printf("maximum: %d\n", (h.IntSlice)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	// Output:
	// maximum: 7
	// 7 3 2 1
}
