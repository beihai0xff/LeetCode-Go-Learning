package Offer

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	if m*n < len(word) {
		return false
	}

	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if k >= len(word) {
			return true
		}
		if i < 0 || j < 0 || i >= m || j >= n || board[i][j] != word[k] {
			return false
		}
		// 将 B 修改（剪枝）为不存在的字符，杜绝往回找成功的可能性。
		board[i][j] = '0'
		res := dfs(i+1, j, k+1) || dfs(i-1, j, k+1) || dfs(i, j+1, k+1) || dfs(i, j-1, k+1)
		board[i][j] = word[k]
		return res
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
