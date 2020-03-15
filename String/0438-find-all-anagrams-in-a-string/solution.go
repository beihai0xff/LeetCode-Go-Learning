/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/3/15 下午11:27
 * Author: beihai
 */

package _438_find_all_anagrams_in_a_string

/*
给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。

说明：
	- 字母异位词指字母相同，但排列不同的字符串。
	- 不考虑答案输出的顺序。
*/

func findAnagrams(s string, p string) []int {
	var freq [256]int
	var result []int
	if len(s) == 0 || len(s) < len(p) {
		return result
	}
	for i := 0; i < len(p); i++ {
		freq[p[i]]++
	}
	left, right, count := 0, 0, len(p)

	for right < len(s) {
		if freq[s[right]] >= 1 {
			count--
		}
		freq[s[right]]--
		right++
		if count == 0 {
			result = append(result, left)
		}
		if right-left == len(p) {
			if freq[s[left]] >= 0 {
				count++
			}
			freq[s[left]]++
			left++
		}

	}
	return result
}
