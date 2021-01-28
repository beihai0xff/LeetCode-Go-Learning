/*
 * Copyright (c) 2021.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2021/1/28 下午9:57
 * Author: beihai
 */

package String

/*
给定两个字符串 s 和 t，它们只包含小写字母。

字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。

请找出在 t 中被添加的字母。
*/

// 位图
func findTheDifference(s string, t string) byte {
	sBit, tBit := make([]byte, 26), make([]byte, 26)
	for i := 0; i < len(s); i++ {
		sBit[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		tBit[t[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		if tBit[i] != sBit[i] {
			return byte('a' + i)
		}
	}
	return 0
}
