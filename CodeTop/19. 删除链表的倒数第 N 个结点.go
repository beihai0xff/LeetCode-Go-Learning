package CodeTop

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast, prev := head, head, &ListNode{Next: head}
	header := prev
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
		prev = prev.Next
	}
	prev.Next = slow.Next
	slow.Next = nil
	return header.Next
}
