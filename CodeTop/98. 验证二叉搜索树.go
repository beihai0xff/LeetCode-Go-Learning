package CodeTop

import "math"

// 中序遍历
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}

	var stack []*TreeNode
	preVal := math.MinInt

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= preVal {
			return false
		}
		preVal = root.Val
		root = root.Right
	}

	return true
}
