/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/12/10 下午9:32
 * Author: beihai
 */

package _206_reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代 aka 双指针
func reverseList0(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var prev *ListNode
	curr := head
	for curr != nil {
		nextNode := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextNode
	}
	return prev
}

// 递归
func reverseList1(head *ListNode) *ListNode {
	// 直接返回最后一个值
	if head == nil || head.Next == nil {
		return head
	}
	temp := reverseList1(head.Next)
	// 让 head 下一个节点的 Next 指向 head ，实现一次局部反转
	head.Next.Next = head
	head.Next = nil
	return temp
}

// 双指针
func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 定义指针 curr，初始化为 head
	curr := head
	for head.Next != nil {
		temp := head.Next.Next
		// 每次都让 head 下一个节点的 Next 指向 curr ，实现一次局部反转
		head.Next.Next = curr
		// 局部反转完成之后，curr 和 head 的 Next 指针同时往前移动一个位置
		curr = head.Next
		head = temp
	}
	return curr
}
