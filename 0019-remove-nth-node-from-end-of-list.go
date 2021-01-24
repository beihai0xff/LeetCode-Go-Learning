/*
 * Copyright (c) 2021.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2021/1/23 下午9:22
 * Author: beihai
 */

package LeetCode_Go_Learing

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	prevHead := new(ListNode)
	prevHead.Next = head
	slow, fast := prevHead, prevHead
	for n > 0 {
		fast = fast.Next
		n--
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	// 此时 slow 与 fast 之间相差 n 步
	// slow 指向 倒数第 n+1 个节点
	slow.Next = slow.Next.Next
	return prevHead.Next
}
