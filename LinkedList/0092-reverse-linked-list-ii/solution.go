/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/12/11 下午9:02
 * Author: beihai
 */

package _092_reverse_linked_list_ii

// 反转链表的中间部分

type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || m >= n {
		return head
	}
	newHead := &ListNode{Val: 0, Next: head}
	pre := newHead
	// 找到第 m 个节点
	for count := 0; pre.Next != nil && count < m-1; count++ {
		pre = pre.Next
	}
	if pre.Next == nil {
		return head
	}
	cur := pre.Next
	// 双指针迭代
	for i := 0; i < n-m; i++ {
		temp := pre.Next
		pre.Next = cur.Next
		cur.Next = cur.Next.Next
		pre.Next.Next = temp
	}
	return newHead.Next
}
