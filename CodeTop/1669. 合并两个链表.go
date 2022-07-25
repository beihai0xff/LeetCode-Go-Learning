package CodeTop

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	dummyHead := &ListNode{Next: list1}
	curr := dummyHead
	for i := 0; i < a; i++ {
		curr = curr.Next
	}
	list1Next := curr.Next
	for i := a; i <= b; i++ {
		list1Next = list1Next.Next
	}

	curr.Next = list2
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = list1Next

	return dummyHead.Next
}
