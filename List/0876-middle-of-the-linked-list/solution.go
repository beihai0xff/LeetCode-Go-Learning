/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/12/11 下午9:38
 * Author: beihai
 */

package _876_middle_of_the_linked_list

// 查找一个链表的中间节点
// 如果链表长度是双数，输出中间节点是中位数后面的那个节点

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针：fast 指针每次移动 2 步，slow 指针每次移动 1 步，当 fast 指针走到终点的时候，slow 指针就是中间节点。
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	length, curr := 0, head
	for curr != nil {
		length++
		curr = curr.Next
	}
	if length%2 == 0 {
		return slow.Next
	}
	return slow
}
