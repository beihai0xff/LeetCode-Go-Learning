/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/24/20, 8:17 PM
 * Author: beihai
 */

package _059_spiral_matrix_ii

/*给定一个正整数 n，生成一个包含 1 到 n×n 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。*/

// 模拟法，类似题解 0054
func generateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}
	if n == 1 {
		return [][]int{{1}}
	}
	top, left, bottom, right := 0, 0, n-1, n-1
	// 二维数组初始化
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	i := 1
	for i <= n*n {
		for j := left; j <= right; j++ {
			result[top][j] = i
			i++
		}
		top++
		for j := top; j <= bottom; j++ {
			result[j][right] = i
			i++
		}
		right--
		for j := right; j >= left; j-- {
			result[bottom][j] = i
			i++
		}
		bottom--
		for j := bottom; j >= top; j-- {
			result[j][left] = i
			i++
		}
		left++
	}
	return result
}
