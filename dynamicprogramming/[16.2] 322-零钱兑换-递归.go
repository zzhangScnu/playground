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
func coinChangeII(coins []int, amount int) int {
	result := math.MaxInt
	memo := make(map[int]int)
	var traverse func(coins []int, remainAmount int) int
	traverse = func(coins []int, remainAmount int) int {
		if remainAmount == 0 {
			return 0
		}
		if remainAmount < 0 {
			return -1
		}
		if ways, ok := memo[remainAmount]; ok {
			return ways
		}
		for _, coin := range coins {
			ways := traverse(coins, remainAmount-coin)
			if ways == -1 {
				continue
			}
			result = min(result, ways+1)
		}
		if result == math.MaxInt {
			memo[remainAmount] = -1
		} else {
			memo[remainAmount] = result
		}
		return memo[remainAmount]
	}
	return traverse(coins, amount)
}

/**
递归版本的动态规划，实际上更像带有备忘录能力的暴力穷举解法。

for _, coin := range coins {
    // result初始化为math.MaxInt，且在递归结束回到本层后，进行择优操作
	ways := traverse(coins, remainAmount-coin)
	if ways == -1 {
		continue
	}
	// 如果无法凑到目标金额，ways == -1的情况下，不会走到这里，即result仍为最大值
	result = min(result, ways+1)
}
// 所以如果result未被更新过，则忘备忘录中写入-1，这样在最终返回的memo[remainAmount]中，
// 才能实现如果没有任何一种硬币组合能组成总金额，返回-1
if result == math.MaxInt {
	memo[remainAmount] = -1
} else {
	memo[remainAmount] = result
}
return memo[remainAmount]
*/
