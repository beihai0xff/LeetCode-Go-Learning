package CodeTop

// 方法一：自顶向下模拟
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBalanced(root.Left) && isBalanced(root.Right) && abs(height(root.Left)-height(root.Right)) <= 1
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(height(root.Left), height(root.Right)) + 1
}

// 方法二：自底向上的递归
func isBalanced2(root *TreeNode) bool {
	return height2(root) >= 0
}

func height2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	if leftHeight == -1 || rightHeight == -1 || abs(leftHeight-rightHeight) > 1 {
		return -1
	}
	return max(leftHeight, rightHeight) + 1
}
