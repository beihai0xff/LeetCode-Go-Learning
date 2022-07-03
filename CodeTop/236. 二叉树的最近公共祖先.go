package CodeTop

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil && right == nil { // 说明 root 的左 / 右子树中都不包含 p,q
		return nil
	}
	if left != nil && right != nil { // 当 left 和 right 同时不为空 ：说明 p,q 分列在 root 的 左/右子），因此 root 为最近公共祖先
		return root
	}
	if left != nil {
		return left
	}
	return right

}
