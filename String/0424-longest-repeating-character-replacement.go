/*
 * Copyright (c) 2021.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2021/1/21 下午9:52
 * Author: beihai
 */

package String

import (
	. "LeetCode/Util"
)

/*
给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，总共可最多替换 k 次。
在执行上述操作后，找到包含重复字母的最长子串的长度。
*/

// 滑动窗口+贪心
func characterReplacement(s string, k int) int {
	count := map[byte]int{} // 每个字符出现的次数
	left, right, maxFreq := 0, 0, 0

	for right < len(s) {
		count[s[right]]++
		maxFreq = Max(maxFreq, count[s[right]]) // 目前出现次数最多的字符
		if right-left+1-maxFreq > k {           // k 不足以将当前窗内所有字符都替换成最高频字符
			count[s[left]]--
			if count[s[left]] < 0 {
				count[s[left]] = 0
			}
			left++ // 左指针前移，先删除最左边的字符
		}
		right++
	}
	return right - left // right 是加过 1 的了
}
