package _138

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	curr, cache := head, map[*Node]*Node{}
	for curr != nil {
		cache[curr] = &Node{Val: curr.Val}
		curr = curr.Next
	}
	dummyHead := &Node{Next: cache[head]}
	curr, newCurr := head, dummyHead.Next
	for curr != nil {
		newCurr.Next, newCurr.Random = cache[curr.Next], cache[curr.Random]
		curr, newCurr = curr.Next, newCurr.Next
	}

	return dummyHead.Next
}
