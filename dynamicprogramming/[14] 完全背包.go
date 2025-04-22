package dynamicprogramming

func KnapsackIII(W int, weights []int, values []int) int {
	n := len(weights)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, W)
	}
	for j := weights[0]; j <= W; j++ {
		dp[0][j] = dp[0][j-weights[0]] + values[0]
	}
	for i := 1; i < n; i++ {
		for j := 1; j <= W; j++ {
			if j < weights[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-weights[i]]+values[i])
			}
		}
	}
	return dp[n-1][W]
}

/**
DP数组及下标含义：
- i：对物品i取/不取，[0, n-1]；
- j：背包容量为j，[0，W]
- dp[i][j]：对物品i取/不取且背包容量为j时，背包中所能装下的物品的最大总价值。

递推公式：
dp[i][j] = max(dp[i-1][j], dp[i][j-weights[i]] + values[i])
不取物品i：dp[i-1][j]
- i-1：只对物品[0, i-1]取/不取；
- j：当前背包容量；
取物品i：dp[i][j-weights[i]] + values[i]
- i：由于是完全背包，物品i可以重复选取，所以依然可以基于i来选取i。若是0/1背包，因为每个物品只能选一次，表达式会是dp[i-1][j-weights[i]] + values[i]；
- j：需要空出j-weights[i]的容量来装下物品i；

遍历方向：
dp[i][j]的值由【左方】及【上方】的已知dp计算结果推导，所以遍历方向是从左往右，由上到下。
（0/1背包则是左上方和上方）
背包和物品的遍历顺序可以颠倒，只要能保证左方和上方的值先被计算出来，则不会影响最终结果。

初始化：
由遍历方向可知，需要初始化dp数组的第一行和第一列。
- 第一列：对于背包容量为0的情况，任何物品都放不下。所以dp[0, n-1][0] = 0；
- 第一行：对于物品0，在背包容量能放下的时候，可以重复选取放入。
*/
