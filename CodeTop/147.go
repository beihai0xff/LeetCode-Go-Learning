package CodeTop

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummyHead, curr, prevNode := &ListNode{Next: head}, head.Next, head
	for curr != nil {
		if prevNode.Val <= curr.Val {
			prevNode = curr
		} else {
			prevNode.Next = curr.Next
			start := dummyHead
			for start.Next.Val <= curr.Val {
				start = start.Next
			}
			start.Next, curr.Next = curr, start.Next
		}
		curr = prevNode.Next
	}

	return dummyHead.Next
}
