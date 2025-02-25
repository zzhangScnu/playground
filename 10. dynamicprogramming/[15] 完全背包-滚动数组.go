package dynamicprogramming

func KnapsackIV(W int, weights []int, values []int) int {
	dp := make([]int, W+1)
	for i := 0; i < len(weights); i++ {
		for j := 0; j <= W; j++ {
			if j >= weights[i] {
				dp[j] = max(dp[j], dp[j-weights[i]]+values[i])
			}
		}
	}
	return dp[W]
}

/**
在dp[i][j] = max(dp[i-1][j], dp[i][j-weights[i]] + values[i])的基础上，将上一层的值拷贝到本层进行计算 ->
dp[i][j] = max(dp[i][j], dp[i][j-weights[i]] + values[i]) -> 直接消去i，用一维数组表达 ->
dp[j] = max(dp[j], dp[j-weights[i]] + values[i])

DP数组及下标含义：
- j：背包容量为j，[0，W]
- dp[j]：对物品i取/不取且背包容量为j时，背包中所能装下的物品的最大总价值。

遍历方向：
对于0/1背包的滚动数组解法，物品和背包的遍历方式如下：
for i := 0; i < n; i++ { // 物品
	for j := W; j >= weights[i]; j-- { // 背包
对背包的遍历，需要从后往前。
【前 -> 后】还是【后 -> 前】，本质上是看计算依赖的前置条件。
如果依赖于上层的dp数组的值，则从前到后；
如果依赖于本层的dp数组的值，则从后到前。
如果先物品再背包 + 背包从前到后遍历，在固定物品i时，对dp[j]的计算依赖于dp[j-weight[i]]的计算，而dp[j-weight[i]又是基于取/不取物品i得来的，会造成对物品i重复选取的可能。
所以对0/1背包，需要【后 -> 前】遍历背包，避免物品重复选取；
而对完全背包，需要【前 -> 后】遍历背包，允许物品重复选取。

再看物品和背包两个循环的顺序：
dp[j]是根据下标j之前所对应的dp[j]计算出来的。只要保证下标j之前的dp[j]都是经过计算的就可以了。
对于0/1背包，因为背包要从后往前遍历，所以只能先遍历物品，再遍历背包。否则下标j之前的dp[j]都是没有经过计算的，导致每个背包中都只能装1个物品；
但对于完全背包，物品和背包的顺序则可以颠倒。

初始化：
dp[j] = 0
*/
