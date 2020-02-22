/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/22/20, 10:33 PM
 * Author: beihai
 */

package _054_spiral_matrix

/*给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。*/

// 按层模拟
func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	if n == 0 {
		return nil
	}
	// top、left、right、bottom 分别是剩余区域的上、左、右、下的下标
	top, left, bottom, right := 0, 0, m-1, n-1
	var result []int

	// 外层循环每次遍历一圈
	for top <= bottom && left <= right {
		// 上
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		// 右
		for i := top + 1; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		if top < bottom && left < right {
			// 下
			for i := right - 1; i > left; i-- {
				result = append(result, matrix[bottom][i])
			}
			// 	左
			for i := bottom; i > top; i-- {
				result = append(result, matrix[i][left])
			}
		}

		// 进入到下一层
		top, left, bottom, right = top+1, left+1, bottom-1, right-1
	}
	return result
}
