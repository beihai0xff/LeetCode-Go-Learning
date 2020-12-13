/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/12/12 下午10:08
 * Author: beihai
 */

package dataStructure

import (
	"container/list"
	"fmt"
)

type BinaryTree struct {
	Val   int
	Left  *BinaryTree
	Right *BinaryTree
}

func (b *BinaryTree) New(val int) *BinaryTree {
	return &BinaryTree{
		Val: val,
	}
}

func (b *BinaryTree) Insert(val int) {
	if val == b.Val {
		return
	} else if val < b.Val {
		if b.Left != nil {
			b.Left.Insert(val)
		} else {
			b.Left = &BinaryTree{
				Val: val,
			}
		}
		return
	} else {
		if b.Right != nil {
			b.Right.Insert(val)
		} else {
			b.Right = &BinaryTree{
				Val: val,
			}
		}
		return
	}
}

func (b *BinaryTree) Search(val int) bool {
	if val == b.Val {
		return true
	} else if val > b.Val {
		return b.Left.Search(val)
	} else {
		return b.Right.Search(val)
	}
}

func (b *BinaryTree) Delete(val int) bool {
	// TODO: 删除的是 root 节点，可以在第一二三种情况中分别处理，但是感觉不优雅
	/*	if  b.Val == val {
		b.Val = 0
	}*/

	current, parent := b, b
	for current != nil {

		if current.Val > val {
			parent = current
			current = current.Left
		} else if current.Val < val {
			parent = current
			current = current.Right
		} else {
			// 第一种情况：该节点是叶节点，直接删除
			if current.Left == nil && current.Right == nil {
				if parent.Left == current {
					parent.Left = nil
				} else {
					parent.Right = nil
				}
				// 删除的是 root 节点
				if b.Val == val {
					b.Val = 0
				}
			} else if current.Left != nil && current.Right == nil {
				// 第二种情况，该节点有左子树或者右子树（不同时有）
				// 直接删除节点，并且该节点的父节点指向其左子节点

				if parent.Val > val {
					parent.Left = current.Left
				} else {
					parent.Right = current.Left
				}
			} else if current.Right != nil && current.Left == nil {
				// 第三种情况，该节点有左子树或者右子树（不同时有）
				// 直接删除节点，并且该节点的父节点指向其右子节点
				if parent.Val > val {
					parent.Left = current.Right
				} else {
					parent.Right = current.Right
				}
			} else {
				// 第四种情况，该节点有左右子树。
				// 1.找到该节点的右子树中的最左子节点（也就是右子树中序遍历的第一个节点）
				// 2.把它的值和要删除的节点的值进行交换
				// 3.然后删除最左子节点

				cTree := current.Right
				cParent := cTree
				for cTree.Left != nil {
					cParent = cTree
					cTree = cTree.Left
				}

				// 最左子节点的值与 current 的值进行交换
				current.Val = cTree.Val

				// 删除 最左子节点
				// 如果 最左子节点 含有右节点，那么让它的父节点直接指向它的子节点
				if cTree.Right != nil {
					cParent.Left = cTree.Right
				} else {
					cParent.Left = nil
				}
			}
			return true
		}
	}
	return false
}

// 二叉树反转
func (b *BinaryTree) Reverse() {
	if b == nil {
		return
	}
	b.Left, b.Right = b.Right, b.Left
	b.Left.Reverse()
	b.Right.Reverse()
}

// 广度优先遍历：层序遍历
func (b *BinaryTree) BreadthFirstSearch() [][]int {
	var res [][]int
	// 双向链表，用作队列
	queue := list.New()
	// 在头部插入节点
	queue.PushFront(b)
	for queue.Len() > 0 {
		var temp []int
		length := queue.Len()
		for i := 0; i < length; i++ {
			// 尾部移出
			node := queue.Remove(queue.Back()).(*BinaryTree)
			temp = append(temp, node.Val)
			if node.Left != nil {
				queue.PushFront(node.Left)
			}
			if node.Right != nil {
				queue.PushFront(node.Right)
			}
		}
		res = append(res, temp)
	}
	return res
}

/*			A
		B		C
	  D	  F
    E    G

先序：考察到一个节点后，即刻输出该节点的值，并继续遍历其左右子树。(根左右)
中序：考察到一个节点后，将其暂存，遍历完左子树后，再输出该节点的值，然后遍历右子树。(左根右)
后序：考察到一个节点后，将其暂存，遍历完左右子树后，再输出该节点的值。(左右根)

广度优先: a -> b -> c -> d -> f -> e -> g

前序遍历: a -> b -> d -> e -> f -> g -> c

中序遍历: e -> d -> b -> g -> f -> a -> c

后序遍历: e -> d -> g -> f -> b -> c -> a*/

// 深度优先遍历：前中后序遍历
func (b *BinaryTree) DepthFirstSearch() {
	fmt.Println(b.Val) // 在这里输出是 前序
	if b.Left != nil {
		b.Left.DepthFirstSearch()
	}
	fmt.Println(b.Val) // 在这里输出是 中序
	if b.Right != nil {
		b.Right.DepthFirstSearch()
	}
	fmt.Println(b.Val) // 在这里输出是 后序
}
