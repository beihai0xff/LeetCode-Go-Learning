package CodeTop

type item struct {
	idx int
	*TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res, queue := 1, []item{{1, root}}
	for len(queue) > 0 {
		if length := queue[len(queue)-1].idx - queue[0].idx + 1; length > res {
			res = length
		}
		var tmp []item
		for _, q := range queue {
			if q.Left != nil {
				tmp = append(tmp, item{q.idx * 2, q.Left})
			}
			if q.Right != nil {
				tmp = append(tmp, item{q.idx*2 + 1, q.Right})
			}
		}
		queue = tmp
	}
	return res
}
