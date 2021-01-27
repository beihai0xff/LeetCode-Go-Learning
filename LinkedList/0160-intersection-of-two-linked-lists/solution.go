/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/12/12 下午3:50
 * Author: beihai
 */

package _160_intersection_of_two_linked_lists

// 寻找两个链表的交叉点，思路类似于链表找环

type ListNode struct {
	Val  int
	Next *ListNode
}

// listA = [0,9,1,2,4], listB = [3,2,4]
// 指针 a 走过的节点：0->9->1->2->4->3->2
// 指针 b 走过的节点：3->2->4->0->9->1->2
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != b {
		if a != nil {
			a = a.Next
		} else {
			a = headB
		}
		if b != nil {
			b = b.Next
		} else {
			b = headA
		}
	}
	return a
}

// 解法 2：分别计算出两个链表的长度，让长的链表先走 n 步，直到两个链表剩余节点长度一样，然后每个链表都走一步，直到节点相等。
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	lenA, lenB := 0, 0
	for a != nil {
		lenA++
		a = a.Next
	}
	for b != nil {
		lenB++
		b = b.Next
	}
	if lenA > lenB {
		for lenA-lenB > 0 {
			headA = headA.Next
			lenA--
		}
	} else {
		for lenB-lenA > 0 {
			headB = headB.Next
			lenB--
		}
	}
	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}
