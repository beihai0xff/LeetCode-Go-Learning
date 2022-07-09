package CodeTop

// 归并排序 + 链表合并 + 求链表长度
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 求链表长度
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}
	dummyHead := &ListNode{Next: head}

	// 归并排序
	for subLength := 1; subLength < length; subLength <<= 1 {
		prev, curr := dummyHead, dummyHead.Next
		for curr != nil {
			// 构造两个需要合并的链表
			head1 := curr
			for i := 1; i < subLength && curr.Next != nil; i++ {
				curr = curr.Next
			}
			head2 := curr.Next
			curr.Next = nil
			curr = head2
			for i := 1; i < subLength && curr != nil && curr.Next != nil; i++ {
				curr = curr.Next
			}

			// 暂时保存剩余的链表节点
			var next *ListNode
			if curr != nil {
				next = curr.Next
				curr.Next = nil
			}
			curr = next

			prev.Next = mergeTwoLists(head1, head2)
			for prev.Next != nil {
				prev = prev.Next
			}

		}
	}

	return dummyHead.Next
}
