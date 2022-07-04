package CodeTop

// 方法一：用线性表存储，构造新链表
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	var queue []*ListNode
	for node := head; node != nil; node = node.Next {
		queue = append(queue, node)
	}
	i, j := 0, len(queue)-1
	for i < j {
		queue[i].Next = queue[j]
		i++
		if i == j {
			break
		}
		queue[j].Next = queue[i]
		j--
	}
	queue[i].Next = nil
}

// 方法二：寻找链表中间节点，反转后半部分链表，两个链表合并
func reorderList2(head *ListNode) {
	if head == nil {
		return
	}
	// 1. 寻找链表中间节点
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	l2 := slow.Next
	slow.Next = nil
	mergeList(head, reverseList(l2))
}

func mergeList(l1, l2 *ListNode) {
	var l1Tmp, l2Tmp *ListNode
	for l1 != nil && l2 != nil {
		l1Tmp = l1.Next
		l2Tmp = l2.Next

		l1.Next = l2
		l2.Next = l1Tmp
		l1 = l1Tmp
		l2 = l2Tmp
	}
}
