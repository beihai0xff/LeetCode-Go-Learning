package CodeTop

// 给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
// 方法一：使用辅助数组，对于矩阵中的第一行而言，在旋转后，它出现在倒数第一列的位置
// 对于矩阵中的第二行而言，在旋转后，它出现在倒数第二列的位置
func rotate1(matrix [][]int) {
	n, tmp := len(matrix), make([][]int, len(matrix))
	for i := range tmp {
		tmp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			tmp[j][n-1-i] = matrix[i][j]
		}
	}

	copy(matrix, tmp)
}

// 方法二：先水平轴翻转，再主对角线翻转
func rotate2(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return
}
