package CodeTop

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	var dfs func(root *TreeNode, curSum int, curPath []int)
	dfs = func(root *TreeNode, currSum int, currPath []int) {
		currPath = append(currPath, root.Val)
		currSum += root.Val

		if root.Left == nil && root.Right == nil {
			if currSum == targetSum {
				tmp := make([]int, len(currPath))
				copy(tmp, currPath)
				res = append(res, tmp)
			}
			return
		}
		if root.Left != nil {
			dfs(root.Left, currSum, currPath)
		}
		if root.Right != nil {
			dfs(root.Right, currSum, currPath)
		}
		currPath = currPath[:len(currPath)-1]
		currSum -= root.Val
	}

	dfs(root, 0, []int{})
	return res
}
