/*
 * Copyright (c) 2021.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2021/1/25 下午9:32
 * Author: beihai
 */

package LinkedList

// 请判断一个链表是否为回文链表。

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	// 找到中间节点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 对后面的节点进行反转
	prev, curr := slow, slow.Next
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	// 从原链表头部和尾部（现为中部）进行遍历对比
	for mid := slow; mid != head; {
		if head.Val == prev.Val {
			prev, head = prev.Next, head.Next
			continue
		} else {
			return false
		}

	}
	return true
}
