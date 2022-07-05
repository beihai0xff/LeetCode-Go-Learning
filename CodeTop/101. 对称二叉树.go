package CodeTop

func isSymmetric(root *TreeNode) bool {
	return checkIsSymmetric(root, root)
}

func checkIsSymmetric(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && checkIsSymmetric(p.Left, q.Right) && checkIsSymmetric(p.Right, q.Left)
}
