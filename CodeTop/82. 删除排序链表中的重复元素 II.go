package CodeTop

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := &ListNode{Next: head}, head
	header := slow
	for fast != nil && fast.Next != nil {
		if fast.Next.Val == fast.Val {
			value := fast.Val
			for fast != nil && fast.Val == value {
				fast = fast.Next
			}
			slow.Next = fast
		} else {
			slow = slow.Next
			fast = fast.Next
		}
	}

	return header.Next
}
