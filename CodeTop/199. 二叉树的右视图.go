package CodeTop

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	queue, res := []*TreeNode{root}, []int{}
	for len(queue) > 0 {
		temp := []*TreeNode{}
		for i := 0; i < len(queue); i++ {
			if queue[i].Left != nil {
				temp = append(temp, queue[i].Left)
			}
			if queue[i].Right != nil {
				temp = append(temp, queue[i].Right)
			}
		}
		res = append(res, queue[len(queue)-1].Val)
		queue = temp
	}

	return res
}
