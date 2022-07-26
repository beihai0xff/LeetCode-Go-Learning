package CodeTop

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		length, tmp := len(queue), []*Node{}
		for i := 0; i < length; i++ {
			node := queue[i]
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
			if i == length-1 {
				node.Next = nil
			} else {
				node.Next = queue[i+1]
			}
		}
		queue = tmp
	}
	return root
}
