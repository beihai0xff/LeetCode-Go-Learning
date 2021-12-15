package Offer

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	prev, curr := head, head.Next
	prev.Next = nil
	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}
	return prev
}
