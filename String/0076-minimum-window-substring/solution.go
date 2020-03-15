/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/3/15 下午10:32
 * Author: beihai
 */

package _076_minimum_window_substring

/*
给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字母的最小子串。

说明：
	- 如果 S 中不存这样的子串，则返回空字符串 ""。
	- 如果 S 中存在这样的子串，我们保证它是唯一的答案。
*/

// 窗口滑动，相似题解：第 3 题，第 76 题，第 438 题，第 567 题
func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return s
	}
	var tFreq, sFreq [256]int
	result, left, right, finalLeft, finalRight, minW, count := "", 0, 0, -1, -1, len(s)+1, 0
	for i := 0; i < len(t); i++ {
		tFreq[t[i]]++
	}
	for left < len(s) {
		if right < len(s) && count < len(t) {
			sFreq[s[right]]++
			if sFreq[s[right]] <= tFreq[s[right]] {
				count++
			}
			right++
		} else {
			if right-left < minW && count == len(t) {
				minW = right - left
				finalLeft = left
				finalRight = right
			}
			if sFreq[s[left]] == tFreq[s[left]] {
				count--
			}
			sFreq[s[left]]--
			left++
		}
	}
	if finalLeft != -1 {
		for i := finalLeft; i < finalRight; i++ {
			result += string(s[i])
		}
	}
	return result
}
