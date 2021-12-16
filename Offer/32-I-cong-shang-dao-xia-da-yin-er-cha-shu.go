package Offer

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue, res := []*TreeNode{root}, []int{}
	for len(queue) != 0 {
		if queue[0].Left != nil {
			queue = append(queue, queue[0].Left)
		}
		if queue[0].Right != nil {
			queue = append(queue, queue[0].Right)
		}
		res = append(res, queue[0].Val)
		queue = queue[1:]
	}
	return res
}
