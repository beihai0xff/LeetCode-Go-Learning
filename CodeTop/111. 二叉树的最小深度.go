package CodeTop

import "math"

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := math.MaxInt
	var dfs func(node *TreeNode, depth int)

	dfs = func(node *TreeNode, depth int) {
		depth++
		if node.Left == nil && node.Right == nil {
			res = min(res, depth)
			return
		}
		if node.Left != nil {
			dfs(node.Left, depth)
		}
		if node.Right != nil {
			dfs(node.Right, depth)
		}
	}

	dfs(root, 0)
	return res
}
