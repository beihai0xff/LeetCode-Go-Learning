package CodeTop

import "math"

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := math.MinInt32
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 递归计算左右子节点的最大贡献值
		// 只有在最大贡献值大于 0 时，才会选取对应子节点
		leftMax, rightMax := max(dfs(node.Left), 0), max(dfs(node.Right), 0)

		// 节点的最大路径和取决于该节点的值与该节点的左右子节点的最大贡献值
		priceNewPath := node.Val + leftMax + rightMax

		res = max(priceNewPath, res)
		return node.Val + max(leftMax, rightMax)
	}

	dfs(root)

	return res
}
