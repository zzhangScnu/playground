package dynamicprogramming

// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
// 示例 1：
//
// 输入：n = 2
// 输出：2
// 解释：有两种方法可以爬到楼顶。
// 1. 1 阶 + 1 阶
// 2. 2 阶
//
// 示例 2：
//
// 输入：n = 3
// 输出：3
// 解释：有三种方法可以爬到楼顶。
// 1. 1 阶 + 1 阶 + 1 阶
// 2. 1 阶 + 2 阶
// 3. 2 阶 + 1 阶
//
// 提示：
//
// 1 <= n <= 45
func climbStairsII(n int) int {
	return climbStairsBySteps(n, 2)
}

func climbStairsBySteps(n int, m int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for j := 1; j <= n; j++ {
		for i := 1; i <= m; i++ {
			if j >= i {
				dp[j] += dp[j-i]
			}
		}
	}
	return dp[n]
}

/**
1. (1, 2)和(2, 1)是不同的爬梯方式，故求的是排列；
2. 题目中n >= 1，故dp[0]无语义，dp[1] = 1，即可以通过走1步到达顶楼，也就是有1种方式可以装满容量为1的背包。
*/
