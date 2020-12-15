/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/12/15 下午8:43
 * Author: beihai
 */

package _738_monotone_increasing_digits

import "strconv"

// 给定一个非负整数 N，找出小于或等于 N 的最大的整数，同时这个整数需要满足其各个位数上的数字是单调递增。

// 从倒数第二位往前遍历，如果前面的值大于后面的值就把当前位数减一，并把后面的值变成 9。
func monotoneIncreasingDigits(N int) int {
	if N < 10 {
		return N
	}
	b := []byte(strconv.Itoa(N))
	for i := len(b) - 2; i > -1; i-- {
		if b[i] > b[i+1] {
			b[i]--
			for j := i + 1; j < len(b); j++ {
				if b[j] == '9' {
					break
				}
				b[j] = '9'
			}
		}
	}
	res, _ := strconv.Atoi(string(b))
	return res
}
