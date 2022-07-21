package CodeTop

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow.Next = reverseList(slow.Next)
	l1, l2 := head, slow.Next
	// len(l1) == len(l2)
	// len(l1) == len(l2) + 1
	for l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l2 = l2.Next
		l1 = l1.Next
	}

	return true
}
