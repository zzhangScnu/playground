package dynamicprogramming

import "math"

// 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
//
// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
//
// 你可以认为每种硬币的数量是无限的。
//
// 示例 1：
//
// 输入：coins = [1, 2, 5], amount = 11
// 输出：3
// 解释：11 = 5 + 5 + 1
//
// 示例 2：
//
// 输入：coins = [2], amount = 3
// 输出：-1
//
// 示例 3：
//
// 输入：coins = [1], amount = 0
// 输出：0
//
// 提示：
//
// 1 <= coins.length <= 12
// 1 <= coins[i] <= 2³¹ - 1
// 0 <= amount <= 10⁴
func coinChange(coins []int, amount int) int {
	maxVal := math.MaxInt64 - 1
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = maxVal
	}
	for i := 0; i < len(coins); i++ {
		for j := 0; j <= amount; j++ {
			if j >= coins[i] {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}
	if dp[amount] == maxVal {
		return -1
	}
	return dp[amount]
}

/**
使用一组硬币候选集去凑一个金额，且每个硬币数量都是无限的，即可以选多次，是典型的完全背包问题。
硬币候选集 -> 物品；
需要凑的金额 -> 背包容量。

DP数组及下标含义：
j：背包容量为j，[0, amount]；
dp[j]：背包容量为j时，能够尽量装满的最少物品个数。即凑成金额j的最少硬币个数。

递推公式：
dp[j] = min(dp[j], dp[j-coins[i]]+1)

初始化：
dp[0] = 0：当背包容量为0时，最少硬币个数为0个；
其余单元格：因为最终要求最小值，所以初始化为最大值。
这里需要减去1，否则在计算dp[j-coins[i]]+1时，如果是dp[j-coins[i]]仍是初始值，+1会导致溢出。

遍历方向：
物品和背包都是从左到右；
因为此处求的是硬币最小值，无关组合/排列，所以两者间的遍历顺序可以颠倒。
*/
