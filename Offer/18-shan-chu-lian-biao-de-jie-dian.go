package Offer

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		next := head.Next
		head.Next = nil
		return next
	}
	prev, curr := head, head.Next
	for curr != nil && curr.Val != val {
		prev, curr = curr, curr.Next
	}
	if curr != nil {
		prev.Next = curr.Next
		curr.Next = nil
	} else {
		prev.Next = nil
	}
	return head
}
