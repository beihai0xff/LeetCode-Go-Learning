/*
 * Copyright (c) 2021.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2021/1/29 下午9:34
 * Author: beihai
 */

package String

import "strconv"

/*
给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
*/

func addStrings(num1 string, num2 string) string {
	add, res := 0, "" // add 是两个数的和的十位
	for len1, len2 := len(num1)-1, len(num2)-1; len1 >= 0 || len2 >= 0 || add > 0; {
		var a, b int
		if len1 >= 0 {
			a = int(num1[len1]) - '0'
			len1--
		}
		if len2 >= 0 {
			b = int(num2[len2]) - '0'
			len2--
		}
		tmp := a + b + add
		res = strconv.Itoa(tmp%10) + res
		add = tmp / 10
	}
	return res
}
