/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/3/9 下午9:14
 * Author: beihai
 */

package _007_reverse_integer

/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
*/

func reverse(x int) int {
	temp := 0
	for x != 0 {
		temp = temp*10 + x%10
		x = x / 10
	}
	// 判断是否超出表示范围，64 位机器下 int 值等效为 int64，题目要求为 int32，不会产出内存溢出问题
	if temp > 1<<31-1 || temp < -(1<<31) {
		return 0
	}
	return temp
}
