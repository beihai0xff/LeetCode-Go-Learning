/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/3/16 下午9:40
 * Author: beihai
 */

package _030_substring_with_concatenation_of_all_words

/*
给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。
*/

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}
	var result = []int{}
	counter := map[string]int{} // 将字符串数组里面的所有字符串都存到 map 中，并累计出现的次数
	for _, w := range words {
		counter[w]++
	}
	length, totalLen, tmpCounter := len(words[0]), len(words[0])*len(words), copyMap(counter)
	// 从源字符串从头开始扫，每次判断字符串数组里面的字符串时候全部都用完了(计数是否为 0)
	for i, start := 0, 0; i < len(s)-length+1 && start < len(s)-length+1; i++ {
		// 如果全部都用完了，并且长度正好是字符串数组任意排列组合的总长度，就记录下这个组合的起始下标
		if tmpCounter[s[i:i+length]] > 0 {
			tmpCounter[s[i:i+length]]--
			if checkWords(tmpCounter) && (i+length-start == totalLen) {
				result = append(result, start)
				continue
			}
			i = i + length - 1
		} else {
			// 如果不符合，就继续考察源字符串的下一个字符，直到扫完整个源字符串。
			start++
			i = start - 1
			tmpCounter = copyMap(counter)
		}
	}
	return result
}

func checkWords(s map[string]int) bool {
	flag := true
	for _, v := range s {
		if v > 0 {
			flag = false
			break
		}
	}
	return flag
}

func copyMap(s map[string]int) map[string]int {
	c := map[string]int{}
	for k, v := range s {
		c[k] = v
	}
	return c
}
