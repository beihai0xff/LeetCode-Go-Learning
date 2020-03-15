/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/3/15 下午8:56
 * Author: beihai
 */

package _003_longest_substring_without_repeating_characters

import (
	. "LeetCode/Util"
	"math/big"
)

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
*/

// 窗口滑动 + 位图，相似题解：第 3 题，第 76 题，第 438 题，第 567 题
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	// 扩展 ASCII 码的位图表示（BitSet），共有 256 位
	var bitSet [256]uint8
	result, left, right := 0, 0, 0
	for left < len(s) {
		// 滑动窗口的右边界不断的右移，只要没有重复的字符，就持续向右扩大窗口边界。
		if right < len(s) && bitSet[s[right]] == 0 {
			// 标记对应的 ASCII 码为 1
			bitSet[s[right]] = 1
			right++
		} else {
			// 出现了重复字符就缩小左边界，直到重复的字符移出了左边界，然后继续移动滑动窗口的右边界。
			// 标记对应的 ASCII 码为 0
			bitSet[s[left]] = 0
			left++
		}
		// 判断是否需要更新最大长度
		result = Max(result, right-left)
	}

	return result
}

// 尝试使用真正的位图--但是更慢了😀
func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}
	var bits big.Int
	result, left, right := 0, 0, 0
	for left < len(s) {
		if right < len(s) && bits.Bit(int(s[right])) == 0 {
			// 标记对应的 ASCII 码为 1
			bits.SetBit(&bits, int(s[right]), 1)
			right++
		} else {
			bits.SetBit(&bits, int(s[left]), 0)
			left++
		}
		// 判断是否需要更新最大长度
		result = Max(result, right-left)
	}

	return result
}
