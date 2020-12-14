/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/12/14 下午2:31
 * Author: beihai
 */

package _345_reverse_vowels_of_a_string

// 反转字符串中的元音字母

func reverseVowels(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; {
		if isVowels(b[i]) && isVowels(b[j]) {
			b[i], b[j] = b[j], b[i]
			i++
			j--
		} else if !isVowels(b[j]) {
			j--
		} else if !isVowels(b[i]) {
			i++
		}
	}
	return string(b)
}

func isVowels(s byte) bool {
	if s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' || s == 'A' || s == 'E' || s == 'I' || s == 'O' || s == 'U' {
		return true
	}
	return false
}
