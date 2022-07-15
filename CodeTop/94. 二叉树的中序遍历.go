package CodeTop

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	for len(stack) > 0 || root != nil {
		if root != nil {
			// 这是模拟递归调用
			// 不断往左子树方向走，每走一次就将当前节点保存到栈中
			stack = append(stack, root)
			root = root.Left
		} else {
			// 当前节点为空，说明左边走到头了，从栈中弹出节点并保存
			// 然后转向右边节点，继续上面整个过程
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, tmp.Val)
			root = tmp.Right
		}

	}
	return res
}
