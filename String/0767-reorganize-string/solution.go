/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/3/16 下午10:28
 * Author: beihai
 */

package _767_reorganize_string

import "sort"

/*
给定一个字符串S，检查是否能重新排布其中的字母，使得两相邻的字符不同。

若可行，输出任意可行的结果。若不可行，返回空字符串。
*/

// 贪心算法：只有当一个字符出现的次数超过 (N+1) / 2 的情况下，这个问题才无解
func reorganizeString(S string) string {
	if S == "" {
		return ""
	}
	sMap := map[byte]int{}
	cMap := map[int][]byte{}
	sByte := []byte(S)
	for _, v := range sByte {
		sMap[v]++
		if sMap[v] > (len(sByte)+1)/2 {
			return ""
		}
	}
	for key, value := range sMap {
		cMap[value] = append(cMap[value], key)
	}

	var keys []int
	for k := range cMap {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	fs := make([]byte, 0)
	for _, k := range keys {
		for i := 0; i < len(cMap[k]); i++ {
			for j := 0; j < k; j++ {
				fs = append(fs, cMap[k][i])
			}
		}
	}
	if fs == nil {
		return ""
	}
	result := ""
	j := (len(fs)-1)/2 + 1
	for i := 0; i <= (len(fs)-1)/2; i++ {
		result += string(fs[i])
		if j < len(fs) {
			result += string(fs[j])
		}
		j++
	}
	return result
}
