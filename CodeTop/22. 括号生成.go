package CodeTop

func generateParenthesis(n int) []string {
	var res []string

	var dfs func(lRemain int, rRemain int, path string)
	dfs = func(lRemain int, rRemain int, path string) {
		if len(path) == 2*n {
			res = append(res, path)
			return
		}

		if lRemain > 0 {
			dfs(lRemain-1, rRemain, path+"(")
		}
		if lRemain < rRemain {
			dfs(lRemain, rRemain-1, path+")")
		}
	}

	dfs(n, n, "")
	return res
}
