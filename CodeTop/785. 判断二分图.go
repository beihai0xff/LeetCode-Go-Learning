package CodeTop

const (
	Uncolored = iota
	Red
	Green
)

var (
	color []int
)

func isBipartite(graph [][]int) bool {
	n := len(graph)
	color = make([]int, n)
	for i := 0; i < n; i++ {
		if color[i] == Uncolored {
			if !dfs(i, Red, graph) {
				return false
			}
		}
	}
	return true
}

func dfs(node, c int, graph [][]int) bool {
	color[node] = c
	nextColor := Red
	if c == Red {
		nextColor = Green
	}
	for _, neighbor := range graph[node] {
		if color[neighbor] == Uncolored {
			if !dfs(neighbor, nextColor, graph) {
				return false
			}
		} else if color[neighbor] != nextColor {
			return false
		}
	}
	return true
}
