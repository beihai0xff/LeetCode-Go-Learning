/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/12/20 下午1:08
 * Author: beihai
 */

package _121_best_time_to_buy_and_sell_stock

/*给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。*/

// 动态规划
func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	buyingCost, maxProfit := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if maxProfit < maxProfit+prices[i]-buyingCost {
			maxProfit += prices[i] - buyingCost
			buyingCost = prices[i]
		}
		if buyingCost > prices[i] {
			buyingCost = prices[i]
		}
	}
	return maxProfit
}

// 优雅的写法
// 逢跌买入，逢涨卖出，连续涨可以用两两相减累加来计算，相当于涨到波峰的最大值减去谷底的值
func maxProfit1(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	maxProfit := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			maxProfit += prices[i+1] - prices[i]
		}
	}
	return maxProfit
}
