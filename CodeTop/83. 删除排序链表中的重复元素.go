package CodeTop

func deleteDuplicatesI(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	curr := head
	for curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return head
}
