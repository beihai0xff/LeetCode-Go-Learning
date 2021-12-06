/*
 * Copyright (c) 2021.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2021/12/6 下午11:36
 * Author: beihai
 */

package Offer

// 请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

var replaced = []byte("%20")

func replaceSpace(s string) string {
	var b []byte
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			b = append(b, replaced...)
		} else {
			b = append(b, s[i])
		}
	}
	return string(b)
}

// 双指针原地修改
func replaceSpace1(s string) string {
	b := []byte(s)
	SpaceNums := 0
	length := len(b)
	for i := 0; i < length; i++ {
		if b[i] == ' ' {
			SpaceNums++
		}
	}
	tmp := make([]byte, SpaceNums*2)
	b = append(b, tmp...)

	for i, j := length-1, len(b)-1; i >= 0; i-- {
		if b[i] == ' ' {
			b[j] = '0'
			b[j-1] = '2'
			b[j-2] = '%'
			j -= 3
		} else {
			b[j] = b[i]
			j--
		}
	}
	return string(b)
}
