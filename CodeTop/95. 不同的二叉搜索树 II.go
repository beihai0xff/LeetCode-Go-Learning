package CodeTop

// 递归 + 回溯
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	var helper func(start, end int) []*TreeNode
	helper = func(start, end int) []*TreeNode {
		if start > end {
			return []*TreeNode{nil}
		}
		allTrees := []*TreeNode{}
		// 枚举可行根节点
		for i := start; i <= end; i++ {
			leftTrees := helper(start, i-1)
			rightTrees := helper(i+1, end)
			for _, left := range leftTrees {
				for _, right := range rightTrees {
					allTrees = append(allTrees, &TreeNode{Val: i, Left: left, Right: right})
				}
			}
		}
		return allTrees
	}
	return helper(1, n)
}
