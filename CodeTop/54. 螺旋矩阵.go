package CodeTop

func spiralOrder(matrix [][]int) []int {
	var res []int
	for len(matrix) != 0 {
		res = append(res, matrix[0]...)
		matrix = roll(matrix[1:])
	}

	return res
}

func roll(matrix [][]int) [][]int {
	rowLen := len(matrix)
	if rowLen == 0 {
		return [][]int{}
	}
	colLen := len(matrix[0])
	newMatrix := make([][]int, colLen)
	for i := range newMatrix {
		newMatrix[i] = make([]int, rowLen)
	}
	col := len(newMatrix) - 1
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			newMatrix[col-j][i] = matrix[i][j]
		}
	}
	return newMatrix
}

func spiralOrder1(matrix [][]int) []int {
	var res []int
	// 每一圈
	for i := 0; i < len(matrix[0])/2+1; i++ {
		// 上
		for j := i; j < len(matrix[0])-i-1 && len(matrix)-1-i >= i; j++ {
			res = append(res, matrix[i][j])
		}
		// 右
		for j := i; j < len(matrix)-i && i <= len(matrix[0])-1-i; j++ {
			res = append(res, matrix[j][len(matrix[0])-i-1])
		}
		// 下
		for j := len(matrix[0]) - i - 2; j > i && len(matrix)-1-i > i; j-- {
			res = append(res, matrix[len(matrix)-1-i][j])
		}
		// 左
		for j := len(matrix) - 1 - i; j > i && i < len(matrix[0])-1-i; j-- {
			res = append(res, matrix[j][i])
		}
	}
	return res
}
