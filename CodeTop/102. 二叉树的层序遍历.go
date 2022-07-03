package CodeTop

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
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
	return res
}
