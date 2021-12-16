package Offer

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	step := 0
	size := len(matrix) * len(matrix[0])
	res := make([]int, 0, size)
	// 定义四个方向的端点
	top, bottom, right, left := 0, len(matrix)-1, len(matrix[0])-1, 0
	for step < size {
		for i := left; i <= right && step < size; i++ {
			res = append(res, matrix[top][i])
			step++
		}
		top++
		for i := top; i <= bottom && step < size; i++ {
			res = append(res, matrix[i][right])
			step++
		}
		right--
		for i := right; i >= left && step < size; i-- {
			res = append(res, matrix[bottom][i])
			step++
		}
		bottom--
		for i := bottom; i >= top && step < size; i-- {
			res = append(res, matrix[i][left])
			step++
		}
		left++
	}
	return res
}
