/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/12/20 下午2:18
 * Author: beihai
 */

package _316_remove_duplicate_letters

// 给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证返回结果的字典序最小

// 单调栈
func removeDuplicateLetters(s string) string {
	// 每个字符出现次数
	var count [26]int
	// 是否在栈中，存在为true
	var exist [26]bool
	// 单调栈
	var stack []rune

	// 统计字符出现次数
	for _, c := range s {
		count[c-'a']++
	}

	for _, c := range s {
		// 栈中已有c，跳过
		if exist[c-'a'] {
			// 同时减少这个字符出现的次数
			count[c-'a']--
			continue
		}

		// 出栈的核心判断要素：
		// 1. 栈里有元素
		// 2. 栈顶元素大于当前元素c
		// 3. 栈顶元素在后续出现
		for len(stack) > 0 && stack[len(stack)-1] > c && count[stack[len(stack)-1]-'a'] > 0 { // 确保栈里面保存的都是比当前元素小的
			// 进入这里说明栈顶元素大于当前元素，所以不符合字典序，要把栈顶元素出栈

			// 标记为栈中不含栈顶元素
			exist[stack[len(stack)-1]-'a'] = false
			// 删除栈顶元素
			stack = stack[:len(stack)-1]
		}

		// 添加新字符
		stack = append(stack, c)
		// 减少该字符出现次数
		count[c-'a']--
		// 标记栈中有该字符
		exist[c-'a'] = true
	}

	return string(stack)
}
