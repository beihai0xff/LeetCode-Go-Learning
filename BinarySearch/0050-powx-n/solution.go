/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/3/20 下午10:34
 * Author: beihai
 */

package _050_powx_n

/*
实现 pow(x, n) ，即计算 x 的 n 次幂函数。
*/

// 递归，二分法：x^2n == x^n * x^n
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	tmp := myPow(x, n/2)
	if n%2 == 0 {
		return tmp * tmp
	}
	return tmp * tmp * x
}
