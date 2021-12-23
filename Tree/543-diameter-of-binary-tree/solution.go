package _43_diameter_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func diameterOfBinaryTree(root *TreeNode) int {
	res = 0

	depth(root)
	return res
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l, r := depth(root.Left), depth(root.Right)

	res = max(res, l+r)
	return max(l, r) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
