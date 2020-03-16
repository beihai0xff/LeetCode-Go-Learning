/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/3/16 下午9:59
 * Author: beihai
 */

package _093_restore_ip_addresses

import "strconv"

/*
给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
*/

// 回溯算法
func restoreIpAddresses(s string) []string {
	if s == "" {
		return []string{}
	}
	result, ip := []string{}, []int{}
	dfs(s, 0, ip, &result)
	return result
}

// 以 0 开头和超过 255 的数字都是非法的。
func dfs(s string, index int, ip []int, result *[]string) {
	if index == len(s) {
		if len(ip) == 4 {
			*result = append(*result, getString(ip))
		}
		return
	}
	if index == 0 {
		num, _ := strconv.Atoi(string(s[0]))
		ip = append(ip, num)
		dfs(s, index+1, ip, result)
	} else {
		num, _ := strconv.Atoi(string(s[index]))
		next := ip[len(ip)-1]*10 + num
		if next <= 255 && ip[len(ip)-1] != 0 {
			ip[len(ip)-1] = next
			dfs(s, index+1, ip, result)
			ip[len(ip)-1] /= 10
		}
		if len(ip) < 4 {
			ip = append(ip, num)
			dfs(s, index+1, ip, result)
			ip = ip[:len(ip)-1]
		}
	}
}

func getString(ip []int) string {
	res := strconv.Itoa(ip[0])
	for i := 1; i < len(ip); i++ {
		res += "." + strconv.Itoa(ip[i])
	}
	return res
}
