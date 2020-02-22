/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/21/20, 12:39 PM
 * Author: beihai
 */

package _048_rotate_image

/*
给定一个 n × n 的二维矩阵表示一个图像，将图像顺时针旋转 90 度。

说明：
	你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。
*/

// 转置加翻转（将原题目增加了返回值，便于测试）
func rotate(matrix [][]int) [][]int {
	row := len(matrix)
	if row <= 0 {
		return matrix
	}
	column := len(matrix[0])
	/*
	 * 对角线翻转
	 * 1 2 3     1 4 7
	 * 4 5 6  => 2 5 8
	 * 7 8 9     3 6 9
	 */
	for i := 0; i < row; i++ {
		for j := i + 1; j < column; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}
	/*
	 * 每一行轴对称翻转
	 * 1 2 3     1 4 7	   7 4 1
	 * 4 5 6  => 2 5 8  => 8 5 2
	 * 7 8 9     3 6 9	   9 6 3
	 */
	halfColumn := column / 2
	for i := 0; i < row; i++ {
		for j := 0; j < halfColumn; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[i][column-j-1]
			matrix[i][column-j-1] = tmp
		}
	}
	return matrix
}
