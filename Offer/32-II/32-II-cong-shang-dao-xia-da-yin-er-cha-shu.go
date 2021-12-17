package _2_II

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	parentQueue := []*TreeNode{root}
	var res [][]int
	for len(parentQueue) != 0 {
		var (
			childQueue []*TreeNode
			tmp        []int
		)
		for _, node := range parentQueue {
			if node == nil {
				continue
			}
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				childQueue = append(childQueue, node.Left)
			}
			if node.Right != nil {
				childQueue = append(childQueue, node.Right)
			}
		}
		parentQueue = childQueue
		res = append(res, tmp)
	}
	return res
}
