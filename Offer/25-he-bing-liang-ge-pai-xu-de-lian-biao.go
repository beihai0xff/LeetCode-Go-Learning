package Offer

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	curr := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			curr.Next = l1
			curr = curr.Next
			l1 = l1.Next
		} else {
			curr.Next = l2
			curr = curr.Next
			l2 = l2.Next
		}
	}

	if l1 != nil {
		curr.Next = l1
	}
	if l2 != nil {
		curr.Next = l2
	}
	return head.Next
}
