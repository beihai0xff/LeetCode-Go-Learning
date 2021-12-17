package _2_III

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	left, queue, res := true, []*TreeNode{root}, [][]int{}

	for len(queue) != 0 {
		childQueue, tmp := []*TreeNode{}, []int{}
		for _, node := range queue {
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				childQueue = append(childQueue, node.Left)
			}
			if node.Right != nil {
				childQueue = append(childQueue, node.Right)
			}
		}

		if !left {
			reverse(tmp)
			left = !left
		}

		res = append(res, tmp)
		queue = childQueue
	}

	return res
}

func reverse(q []int) {
	if len(q) == 0 {
		return
	}
	left, right := 0, len(q)-1
	for left < right {
		q[left], q[right] = q[right], q[left]
		left++
		right--
	}

	return
}
