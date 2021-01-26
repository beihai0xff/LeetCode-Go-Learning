/*
 * Copyright (c) 2021.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2021/1/26 下午10:01
 * Author: beihai
 */

package Array

/*
给你一个由一些多米诺骨牌组成的列表 dominoes。

如果其中某一张多米诺骨牌可以通过旋转 0 度或 180 度得到另一张多米诺骨牌，我们就认为这两张牌是等价的。

形式上，dominoes[i] = [a, b] 和 dominoes[j] = [c, d] 等价的前提是 a==c 且 b==d，或是 a==d 且 b==c。

在 0 <= i < j < dominoes.length 的前提下，找出满足 dominoes[i] 和 dominoes[j] 等价的骨牌对 (i, j) 的数量。
*/

// 把有序整数拼成一个二位数
func numEquivDominoPairs(dominoes [][]int) (ans int) {
	cnt := [100]int{}
	for _, d := range dominoes {
		// 为了使得「等价」更易于比较，让较小的数排在前面
		if d[0] > d[1] {
			d[0], d[1] = d[1], d[0]
		}
		// 拼数字
		v := d[0]*10 + d[1]
		ans += cnt[v]
		cnt[v]++
	}
	return
}
