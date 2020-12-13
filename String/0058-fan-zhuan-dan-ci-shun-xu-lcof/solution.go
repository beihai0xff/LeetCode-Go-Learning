/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/12/13 下午8:35
 * Author: beihai
 */

package _058_fan_zhuan_dan_ci_shun_xu_lcof

import "strings"

/*输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。
为简单起见，标点符号和普通字母一样处理。
例如输入字符串"I am a student. "，则输出"student. a am I"。*/

// 双指针，倒叙遍历
func reverseWords(s string) string {
	fast, slow := len(s)-1, len(s)-1
	for fast >= 0 && string(s[fast]) == " " {
		fast--
		slow--
	}
	var res string
	for fast >= 0 {
		for fast > 0 && string(s[fast-1]) != " " {
			fast--
		}
		res += s[fast:slow+1] + " " // 注意左闭右开
		for fast > 0 && string(s[fast-1]) == " " {
			fast--
		}
		fast--
		slow = fast
	}
	return res[:len(res)-1]
}

// 空格切割字符串
func reverseWords1(s string) string {
	strList := strings.Split(s, " ")
	var res []string
	for i := len(strList) - 1; i >= 0; i-- {
		str := strings.TrimSpace(strList[i])
		if len(str) > 0 {
			res = append(res, strList[i])
		}
	}
	return strings.Join(res, " ")
}
