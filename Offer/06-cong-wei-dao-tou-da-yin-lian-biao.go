/*
 * Copyright (c) 2021.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2021/12/7 下午12:21
 * Author: beihai
 */

package Offer

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归栈
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	return append(reversePrint(head.Next), head.Val)
}
