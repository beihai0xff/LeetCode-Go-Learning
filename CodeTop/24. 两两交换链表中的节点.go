package CodeTop

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	one, two, three := head, head.Next, head.Next.Next
	two.Next = one
	one.Next = swapPairs(three)
	return two
}
