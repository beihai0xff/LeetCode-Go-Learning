package CodeTop

// 方法一：将有序链表转成有序数组，数组的中点为 root 节点
func sortedListToBST1(head *ListNode) *TreeNode {
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return buildBST(arr, 0, len(arr)-1)
}

func buildBST(arr []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := (start + end) >> 1
	root := &TreeNode{Val: arr[mid]}
	root.Left = buildBST(arr, start, mid-1)
	root.Right = buildBST(arr, mid+1, end)
	return root
}

// 方法二：快慢指针寻找链表中点，链表的中点为 root 节点
func sortedListToBST2(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	var preSlow *ListNode = nil
	for fast != nil && fast.Next != nil {
		preSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	root := &TreeNode{Val: slow.Val}
	if preSlow != nil {
		preSlow.Next = nil
		root.Left = sortedListToBST(head)
	}
	root.Right = sortedListToBST(slow.Next)
	return root
}

var h *ListNode

// 方法三：中序遍历
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	length := 0
	curr := head
	for curr != nil {
		length++
		curr = curr.Next
	}
	h = head
	return buildBST3(0, length-1)
}
func buildBST3(start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := (start + end) >> 1
	left := buildBST3(start, mid-1)
	root := &TreeNode{Val: h.Val}
	h = h.Next
	root.Left = left
	root.Right = buildBST3(mid+1, end)
	return root
}
