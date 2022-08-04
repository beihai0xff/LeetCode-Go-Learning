package CodeTop

import "sort"

// 排序
func groupAnagrams(strs []string) [][]string {
	if strs == nil || len(strs) == 0 {
		return nil
	}

	cache := map[string][]string{}
	for _, str := range strs {
		b := []byte(str)
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		sortedStr := string(b)
		cache[sortedStr] = append(cache[sortedStr], str)
	}

	var res [][]string
	for _, v := range cache {
		res = append(res, v)
	}

	return res
}

// 计数
func groupAnagrams2(strs []string) [][]string {
	if strs == nil || len(strs) == 0 {
		return nil
	}

	cache := map[[26]int][]string{}
	for _, str := range strs {
		count := [26]int{}
		for _, char := range str {
			count[char-'a']++
		}
		cache[count] = append(cache[count], str)
	}

	var res [][]string
	for _, v := range cache {
		res = append(res, v)
	}

	return res
}
