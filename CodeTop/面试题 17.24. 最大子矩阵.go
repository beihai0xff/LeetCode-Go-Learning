package CodeTop

import "math"

func getMaxMatrix(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	res := make([]int, 4)

	max := matrix[0][0]

	for beginLine := 0; beginLine < m; beginLine++ {
		sum := make([]int, n)
		for i := beginLine; i < m; i++ {
			beginRow := 0
			dp := math.MinInt
			// 求连续子数组的最大和
			for j := 0; j < n; j++ {
				sum[j] += matrix[i][j]
				if dp > 0 {
					dp += sum[j]
				} else {
					dp = sum[j]
					beginRow = j
				}
				if max < dp {
					res[0] = beginLine
					res[1] = beginRow
					res[2] = i
					res[3] = j
					max = dp
				}

			}
		}
	}
	return res
}
