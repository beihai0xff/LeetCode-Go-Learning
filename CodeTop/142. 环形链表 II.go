package CodeTop

func detectCycle(head *ListNode) *ListNode {
	cache, slow := map[*ListNode]struct{}{}, head
	for slow != nil {
		if _, ok := cache[slow]; ok {
			return slow
		}
		cache[slow] = struct{}{}
		slow = slow.Next
	}
	return nil
}

func detectCycle2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			for n := head; n != slow; {
				n = n.Next
				slow = slow.Next
			}
			return slow
		}
	}

	return nil
}
