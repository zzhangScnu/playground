package dynamicprogramming

// 给你一个整数数组 prices 和一个整数 k ，其中 prices[i] 是某支给定的股票在第 i 天的价格。
//
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。也就是说，你最多可以买 k 次，卖 k 次。
//
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
// 示例 1：
//
// 输入：k = 2, prices = [2,4,1]
// 输出：2
// 解释：在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
//
// 示例 2：
//
// 输入：k = 2, prices = [3,2,6,5,0,3]
// 输出：7
// 解释：在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
//
//	随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3
//
// 提示：
//
// 1 <= k <= 100
// 1 <= prices.length <= 1000
// 0 <= prices[i] <= 1000
func maxProfitIVI(k int, prices []int) int {
	n := len(prices)
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	for j := 1; j <= k; j++ {
		dp[0][j][0] = -prices[0]
	}
	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j-1][1]-prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j][0]+prices[i])
		}
	}
	return dp[len(prices)-1][k][1]
}

/**
因为最多可以买卖k次，所以可以把k抽象出来作为一个变量承载。

DP数组及下标含义：
- i：第i天；
- j：第j次交易。包含持有/不持有两种状态，[0, k]；
  其实是交易次数。即j ==0时表示没有发生交易。这是为了兼容递推公式中dp[i-1][j-1][1]。
  两次买卖没有这个问题是因为，第一次买入时基于初始本金0，而不是抽象成表达式。
- dp[i][j][0]：第i天，第j次交易过程中，持有股票的最大利润；
- dp[i][j][1]：第i天，第j次交易过程中，不持有股票的最大利润。

递推公式：
dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j-1][1] - prices[i])
若操作买入，需基于上一次交易(j-1)和持有状态1(不持有)；

dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j][0] + prices[i])
若操作卖出，需基于本次交易(j)和持有状态0(持有)。

res = dp[len(prices)-1][k-1][1]

初始化：
可以看出依赖于i-1和j-1来推导出i和j，故需初始化：
i == 0 -> 第0天。需初始化在第0天发生交易的情况。
j == 0 -> 第0次交易，即无交易；
j == 1 -> 第1次交易，买入时利润 == -prices[0]，卖出时利润 == 0；
j == 2 -> 第2次交易，买入时利润 == -prices[0]，卖出时利润 == 0；
……

遍历方向：
从左到右。
*/
