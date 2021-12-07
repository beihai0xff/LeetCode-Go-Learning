/*
 * Copyright (c) 2021.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2021/12/7 下午12:28
 * Author: beihai
 */

package Offer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootIdx := 0
	// 找到中序遍历根节点的位置
	for i := range inorder {
		if inorder[i] == preorder[0] {
			rootIdx = i
			break
		}
	}
	root := &TreeNode{Val: preorder[0]}
	root.Left = buildTree(preorder[1:len(inorder[:rootIdx])+1], inorder[:rootIdx])
	root.Right = buildTree(preorder[len(inorder[:rootIdx])+1:], inorder[rootIdx+1:])
	return root
}
