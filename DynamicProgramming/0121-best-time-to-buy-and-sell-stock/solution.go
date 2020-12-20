/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/12/20 下午1:08
 * Author: beihai
 */

package _121_best_time_to_buy_and_sell_stock

/*给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。

注意：你不能在买入股票前卖出股票。
*/

// 动态规划
func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	buyingCost, maxProfit := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if maxProfit < prices[i]-buyingCost {
			maxProfit = prices[i] - buyingCost
		}
		if buyingCost > prices[i] {
			buyingCost = prices[i]
		}
	}
	return maxProfit
}
