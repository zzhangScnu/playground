package dynamicprogramming

// 给定一个整数数组
// prices，其中第 prices[i] 表示第 i 天的股票价格 。
//
// 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
//
// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
//
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
// 示例 1:
//
// 输入: prices = [1,2,3,0,2]
// 输出: 3
// 解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
//
// 示例 2:
//
// 输入: prices = [1]
// 输出: 0
//
// 提示：
//
// 1 <= prices.length <= 5000
// 0 <= prices[i] <= 1000
func maxProfitVI(prices []int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 4)
	}
	dp[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i], dp[i-1][3]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][3])
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][3] = dp[i-1][2]
	}
	return max(dp[n-1][1], dp[n-1][2], dp[n-1][3])
}

/**
在买卖股票的最佳时机II的基础上稍作调整。

将第i天的持股状态做拆分：
0：持股(买入 / 保持持股状态)
1：保持不持股状态
2：卖出
3：冷静期

递推公式：
dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i], dp[i-1][3]-prices[i])
	持股状态可分两种情况：
	1. 第i天保持持股状态：即延续i-1天的持股状态，从dp[i-1][0]推导而来；
	2. 第i天操作买入：此时i-1天又分为：
		- 保持不持股状态：dp[i-1][1]；
		- 冷静期：dp[i-1][3]；
		只有当i-1天为这两种情况才能在第i天操作买入。
dp[i][1] = max(dp[i-1][1], dp[i-1][3])
	保持不持股状态可分为两种情况：
	1. 第i-1天不持股：dp[i-1][1]
	2. 第i-1天为冷静期：dp[i-1][3]
dp[i][2] = dp[i-1][0] + prices[i]
	第i天卖出只能由第i-1天持股转移而来：dp[i-1][0]
dp[i][3] = dp[i-1][2]
	第i天为冷静期只能由第i-1天操作卖出转移而来：dp[i-1][2]
res = max(dp[len(prices)-1][1], dp[len(prices)-1][2], dp[len(prices)-1][3])

初始化：
i依赖于i-1，故：
dp[0][0] = -prices[0]
dp[0][1] = 0
dp[0][2] = 0
dp[0][3] = 0

遍历方向：
从左到右
*/
