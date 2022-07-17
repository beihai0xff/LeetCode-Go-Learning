package CodeTop

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	for len(lists) > 1 {
		var tmp []*ListNode
		for i := 0; i < len(lists)-1; i += 2 {
			tmp = append(tmp, mergeTwoLists(lists[i], lists[i+1]))
		}
		if len(lists)%2 == 1 {
			tmp = append(tmp, lists[len(lists)-1])
		}
		lists = tmp
	}
	return lists[0]
}
