package CodeTop

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	curr, carry := &ListNode{}, 0
	header := curr
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carry
		curr.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		curr = curr.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		sum := l1.Val + carry
		curr.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		curr = curr.Next
		l1 = l1.Next
	}
	for l2 != nil {
		sum := l2.Val + carry
		curr.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		curr = curr.Next
		l2 = l2.Next
	}
	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}

	return header.Next
}
