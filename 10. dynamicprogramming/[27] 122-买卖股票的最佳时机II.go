package dynamicprogramming

// 给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。
//
// 在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。
//
// 返回 你能获得的 最大 利润 。
//
// 示例 1：
//
// 输入：prices = [7,1,5,3,6,4]
// 输出：7
// 解释：在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4。
// 随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6 - 3 = 3。
// 最大总利润为 4 + 3 = 7 。
//
// 示例 2：
//
// 输入：prices = [1,2,3,4,5]
// 输出：4
// 解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4。
// 最大总利润为 4 。
//
// 示例 3：
//
// 输入：prices = [7,6,4,3,1]
// 输出：0
// 解释：在这种情况下, 交易无法获得正利润，所以不参与交易可以获得最大利润，最大利润为 0。
//
// 提示：
//
// 1 <= prices.length <= 3 * 10⁴
// 0 <= prices[i] <= 10⁴
func maxProfitII(prices []int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0], dp[0][1] = 0, -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][1]+prices[i], dp[i-1][0])
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
	}
	return dp[n-1][0]
}

/**
跟买卖股票I的区别是，可以重复买入和卖出。

第i天持有：
dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
- dp[i-1][1]：保持持有状态，不做操作；
- dp[i-1][0]-prices[i]：第i天买入，将第i-1天的利润减去第i天的股票价格。持有状态从0转移到1；
  基于第i-1天的利润买入股票，即可以多次买入。
  但同时也是基于第i-1天的持有状态0买入股票，即限制了需要先卖出，再买入。保证了同一时刻最多只能持有一支股票。

第i天不持有：
dp[i][0] = max(dp[i-1][1]+prices[i], dp[i-1][0])
同理，第i天卖出的操作基于第i-1天的持有状态1。也就保证了需先买入才能卖出。
*/
