package CodeTop

// 先将 curr 的下一个节点记录为 next；
// 执行操作 ①：把 curr 的下一个节点指向 next 的下一个节点；
// 执行操作 ②：把 next 的下一个节点指向 pre 的下一个节点；
// 执行操作 ③：把 pre 的下一个节点指向 next。

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{Next: head}
	pre := dummyNode
	// 找到 pre 节点
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	curr := pre.Next
	for i := 0; i < right-left; i++ {
		next := curr.Next
		curr.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummyNode.Next
}
