package CodeTop

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return false
	}

	var dfs func(root *TreeNode, currNodes, total int) bool
	dfs = func(root *TreeNode, currNodes, total int) bool {
		if root == nil {
			return true
		}

		if currNodes > total {
			return false
		}

		return dfs(root.Left, currNodes*2, total) && dfs(root.Right, currNodes*2+1, total)
	}

	return dfs(root, 1, countNodes(root))
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
