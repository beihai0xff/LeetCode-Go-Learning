/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/3/15 下午11:23
 * Author: beihai
 */

package _567_permutation_in_string

/*
给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
换句话说，第一个字符串的排列之一是第二个字符串的子串。

注意：
	- 输入的字符串只包含小写字母
	- 两个字符串的长度都在 [1, 10,000] 之间
*/

// 窗口滑动，相似题解：第 3 题，第 76 题，第 438 题，第 567 题
func checkInclusion(s1 string, s2 string) bool {
	var freq [256]int
	if len(s2) == 0 || len(s2) < len(s1) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		freq[s1[i]-'a']++
	}
	left, right, count := 0, 0, len(s1)

	for right < len(s2) {
		if freq[s2[right]-'a'] >= 1 {
			count--
		}
		freq[s2[right]-'a']--
		right++
		if count == 0 {
			return true
		}
		if right-left == len(s1) {
			if freq[s2[left]-'a'] >= 0 {
				count++
			}
			freq[s2[left]-'a']++
			left++
		}
	}
	return false
}
