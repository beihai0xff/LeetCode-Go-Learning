package CodeTop

func kthSmallest(root *TreeNode, k int) int {
	var res int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			tmp := stack[len(stack)-1]
			if k == 1 {
				return tmp.Val
			}
			stack = stack[:len(stack)-1]
			k--
			root = tmp.Right
		}
	}

	return res
}
