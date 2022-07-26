package CodeTop

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var res int
	var dfs func(root *TreeNode, sum int)
	dfs = func(root *TreeNode, sum int) {
		if root.Left == nil && root.Right == nil {
			res += sum*10 + root.Val
		}
		sum = sum*10 + root.Val
		if root.Left != nil {
			dfs(root.Left, sum)
		}
		if root.Right != nil {
			dfs(root.Right, sum)
		}
	}

	dfs(root, 0)
	return res
}
