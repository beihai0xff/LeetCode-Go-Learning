/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/3/23 下午9:49
 * Author: beihai
 */

package _069_sqrtx

/*
实现 int sqrt(int x) 函数。

计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
*/

// 从几何角度来看，整数平方根就是小于 xx 的最大正方形的边长。
// 解法一：二分法
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	left, right, res := 1, x, 0
	for left <= right {
		mid := left + ((right - left) >> 1)
		if mid < x/mid {
			left = mid + 1
			res = mid
		} else if mid == x/mid {
			return mid
		} else {
			right = mid - 1
		}
	}
	return res
}

// 解法二：牛顿迭代法
func mySqrt1(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) >> 1
	}
	return r
}
