package CodeTop

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	var res []string
	var dfs func(*TreeNode, string)
	dfs = func(node *TreeNode, path string) {
		path += strconv.Itoa(node.Val) + "->"
		if node.Left == nil && node.Right == nil {
			res = append(res, path[:len(path)-2])
		}
		if node.Left != nil {
			dfs(node.Left, path)
		}
		if node.Right != nil {
			dfs(node.Right, path)
		}
	}
	dfs(root, "")
	return res
}
