/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/12/17 下午9:18
 * Author: beihai
 */

package _714_best_time_to_buy_and_sell_stock_with_transaction_fee

import "LeetCode/Util"

/*给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格；非负整数 fee 代表了交易股票的手续费用。
你可以无限次地完成交易，但是你每次交易都需要付手续费。
如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。要求返回获得利润的最大值。*/

// 动态规划
// buy 表示前 i 天买入的最大收益，sell 表示前 i 天卖出的最大收益
func maxProfit(prices []int, fee int) int {
	sell, buy := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		sell = Util.Max(sell, buy+prices[i]-fee)
		buy = Util.Max(buy, sell-prices[i])
	}
	return sell
}
