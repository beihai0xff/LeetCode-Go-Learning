package CodeTop

func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}

	var stack, queue []*TreeNode
	for root != nil || len(stack) > 0 {
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
			queue = append(queue, tmp)
			root = tmp.Right
		}
	}

	// 找到位置
	index1, index2 := -1, -1
	for i := 0; i < len(queue)-1; i++ {
		if queue[i].Val > queue[i+1].Val {
			index2 = i + 1
			if index1 == -1 {
				index1 = i
			} else {
				break
			}
		}
	}

	queue[index1].Val, queue[index2].Val = queue[index2].Val, queue[index1].Val
}
