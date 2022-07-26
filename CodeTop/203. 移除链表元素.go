package CodeTop

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{Next: head}
	prev, curr := dummyHead, dummyHead.Next
	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
		curr = curr.Next
	}

	return dummyHead.Next
}
