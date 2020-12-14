/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/12/11 下午9:02
 * Author: beihai
 */

package _021_merge_two_sorted_lists

// 合并两个有序链表

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

// 迭代
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	preHead := ListNode{-1, nil}
	prev := &preHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}
	// l1 或 l2 的最后一个节点还没有添加进来，这里判断一下
	if l1 != nil {
		prev.Next = l1
	} else {
		prev.Next = l2
	}

	return preHead.Next
}
