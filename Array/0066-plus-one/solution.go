/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/25/20, 5:03 PM
 * Author: beihai
 */

package _066_plus_one

/*
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。
*/

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{}
	}
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+carry > 9 {
			digits[i] = 0
			carry = 1
		} else {
			// 没有进位
			digits[i] += carry
			carry = 0
			break
		}
	}
	// 如果最高位有进位需要在数组里面第 0 位插入一个 1
	if digits[0] == 0 && carry == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
