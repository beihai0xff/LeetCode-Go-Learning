/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/12/20 下午1:08
 * Author: beihai
 */

package _121_best_time_to_buy_and_sell_stock

import (
	"LeetCode/Util"
	"math"
)

/*
给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格。

设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:

- 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
- 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
*/

// 动态规划
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	buy, sell := make([]int, len(prices)), make([]int, len(prices))
	// 赋初值，因为可能为负数
	for i := range buy {
		buy[i] = math.MinInt64
	}
	buy[0] = -prices[0]
	buy[1] = Util.Max(buy[0], -prices[1])
	sell[1] = Util.Max(sell[0], buy[0]+prices[1])
	for i := 2; i < len(prices); i++ {
		sell[i] = Util.Max(sell[i-1], buy[i-1]+prices[i])
		buy[i] = Util.Max(buy[i-1], sell[i-2]-prices[i])
	}
	return sell[len(sell)-1]
}
