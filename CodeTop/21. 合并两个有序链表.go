package CodeTop

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	curr := &ListNode{}
	header := curr
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	if list1 != nil {
		curr.Next = list1
	}
	if list2 != nil {
		curr.Next = list2
	}

	return header.Next
}
