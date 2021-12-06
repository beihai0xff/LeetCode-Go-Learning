/*
 * Copyright (c) 2021.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2021/12/6 下午11:29
 * Author: beihai
 */

package Offer

/*
选左上角，往右走和往下走都增大，不能选

选右下角，往上走和往左走都减小，不能选

选左下角，往右走增大，往上走减小，可选

选右上角，往下走增大，往左走减小，可选
*/
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	i, j := len(matrix)-1, 0
	for i >= 0 && j < len(matrix[0]) {
		if matrix[i][j] > target {
			i--
			continue
		}
		if matrix[i][j] < target {
			j++
			continue
		}
		return true
	}
	return false
}
