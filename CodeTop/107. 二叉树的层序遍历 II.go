package CodeTop

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	var res [][]int
	// i 表示第 n 层
	for i := 0; len(queue) > 0; i++ {
		var tmp []*TreeNode
		var tmpRes []int

		for j := 0; j < len(queue); j++ {
			node := queue[j]
			tmpRes = append(tmpRes, node.Val)
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		res = append(res, tmpRes)
		queue = tmp
	}

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	return res
}
