/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/12/12 下午10:08
 * Author: beihai
 */

package dataStructure

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
			} else if current.Left != nil && current.Right == nil {
				// 第二种情况，该节点有左子树或者右子树（不同时有）
				// 直接删除节点，并且该节点的父节点指向其子节点
				if parent.Val > val {
					parent.Left = current.Left
				} else {
					parent.Right = current.Left
				}
				current = nil
			} else if current.Right != nil && current.Left == nil {
				if parent.Val > val {
					parent.Left = current.Right
				} else {
					parent.Right = current.Right
				}
				current = nil
			} else {
				// 第三种情况，该节点有左右子树。将该节点与其右子树看成一棵线索二叉树
				// 找到该节点的直接后继，用该后继节点占领该节点，就完成了删除操作
				// 对右子树（不包括删除的那个节点）进行中序遍历找到直接后继节点

				cltree := current.Right
				clparent := cltree
				for cltree.Left != nil {
					clparent = cltree
					cltree = cltree.Left
				}

				// 用 cltree 占领 current 的位置
				if parent.Val > val {
					parent.Left = cltree

				} else {
					parent.Right = cltree
				}

				// 很显然，cltree是没有左子树的
				cltree.Left = current.Left
				if current.Right != cltree {
					// 说明直接后继不是 current 的右子节点
					// 此时要把 cltree 与其父节点的关系切断，否则会造成死循环
					clparent.Left = nil
					cltree.Right = current.Right
				}
				current = nil
				return true
			}
		}
	}
	return false
}
