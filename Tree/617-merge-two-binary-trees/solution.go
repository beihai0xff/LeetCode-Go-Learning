package _17_merge_two_binary_trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root2 == nil && root1 == nil {
		return nil
	}
	res := &TreeNode{}

	if root2 == nil {
		res = root1
	} else if root1 == nil {
		res = root2
	} else {
		res.Val = root2.Val + root1.Val
		res.Left = mergeTrees(root1.Left, root2.Left)
		res.Right = mergeTrees(root1.Right, root2.Right)
	}

	return res
}
