package Offer

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	}
	if A == nil || B == nil {
		return false
	}

	return helper(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

// helper 校验 B 是否与 A 的一个子树拥有相同的结构和节点值
func helper(a, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	// a.Val == b.Val 递归校验 A B 左子树和右子树的结构和节点是否相同
	return helper(a.Left, b.Left) && helper(a.Right, b.Right)
}
