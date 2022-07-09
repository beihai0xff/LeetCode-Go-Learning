package CodeTop

func rotateRight(matrix [][]int) {
	n := len(matrix)
	// 水平翻转
	for i := 0; i < len(matrix)/2; i++ {
		matrix[i], matrix[len(matrix)-1-i] = matrix[len(matrix)-1-i], matrix[i]
	}
	// 主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func rotateLeft(matrix [][]int) {
	n := len(matrix)
	// 垂直翻转
	m := len(matrix[0])
	for i := 0; i < m/2; i++ {
		for j := 0; j < n; j++ {
			matrix[j][i], matrix[m-1-j][i] = matrix[m-1-j][i], matrix[j][i]
		}

	}
	// 主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
