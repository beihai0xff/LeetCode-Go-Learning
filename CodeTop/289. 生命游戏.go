package CodeTop

// 添加额外状态：活细胞变死细胞 10 -> 1 变 0
// 添加额外状态：死细胞变活细胞 -1 -> 0 变 1
func gameOfLife(board [][]int) {
	if board == nil || len(board) == 0 || len(board[0]) == 0 {
		return
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			aliveNum := isAlive(i-1, j-1, board) + isAlive(i-1, j, board) + isAlive(i-1, j+1, board) + isAlive(i, j-1, board) + isAlive(i, j+1, board) + isAlive(i+1, j-1, board) + isAlive(i+1, j, board) + isAlive(i+1, j+1, board)
			if board[i][j] == 1 {
				if aliveNum < 2 || aliveNum > 3 {
					board[i][j] = 10
				}
			} else if board[i][j] == 0 && aliveNum == 3 {
				board[i][j] = -1
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 10 {
				board[i][j] = 0
			} else if board[i][j] == -1 {
				board[i][j] = 1
			}
		}
	}
}

func isAlive(i, j int, board [][]int) int {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return 0
	}

	if board[i][j] >= 1 {
		return 1
	}
	return 0
}
