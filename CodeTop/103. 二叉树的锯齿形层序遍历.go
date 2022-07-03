package CodeTop

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue, isLeft := []*TreeNode{root}, true
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

		if !isLeft {
			reverse(tmpRes)
		}
		res = append(res, tmpRes)
		queue, isLeft = tmp, !isLeft
	}
	return res
}

func reverse(tmp []int) {
	for i := 0; i < len(tmp)/2; i++ {
		tmp[i], tmp[len(tmp)-1-i] = tmp[len(tmp)-1-i], tmp[i]
	}
}
