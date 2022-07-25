package CodeTop

func rotateRightI(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	length, curr := 1, head
	for curr.Next != nil {
		length++
		curr = curr.Next
	}
	curr.Next = head

	for pos := length - k%length; pos > 0; pos-- {
		curr = curr.Next
	}
	res := curr.Next
	curr.Next = nil

	return res
}
