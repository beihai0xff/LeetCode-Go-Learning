package CodeTop

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	for len(lists) > 1 {
		l1, l2 := lists[0], lists[1]
		lists = lists[2:]
		lists = append(lists, mergeTwoLists(l1, l2))
	}
	return lists[0]
}
