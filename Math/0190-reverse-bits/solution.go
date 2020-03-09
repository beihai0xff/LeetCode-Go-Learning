/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/3/9 下午9:27
 * Author: beihai
 */

package _190_reverse_bits

/*
颠倒给定的 32 位无符号整数的二进制位。
*/

// 对半拆开，互相换位
func reverseBits(num uint32) uint32 {
	num = (num >> 16) | (num << 16)
	num = ((num & 0xff00ff00) >> 8) | ((num & 0x00ff00ff) << 8)
	num = ((num & 0xf0f0f0f0) >> 4) | ((num & 0x0f0f0f0f) << 4)
	num = ((num & 0xcccccccc) >> 2) | ((num & 0x33333333) << 2)
	num = ((num & 0xaaaaaaaa) >> 1) | ((num & 0x55555555) << 1)
	return num
}

// 取当前 n 的最后一位：n & 1
// 将最后一位移动到对应位置，result 左移，num 右移
func reverseBits1(num uint32) uint32 {
	var result uint32
	for i := 0; i < 32; i++ {
		result = result<<1 | num&1
		num >>= 1
	}
	return result
}
