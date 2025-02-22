package dynamicprogramming

/**
有一个容量为W的背包和n个物品，每个物品有两个属性:重量Wi和价值vi(i=1，2，...)
对于每个物品，你只能选择放入背包或者不放入背包(即"0-1"选择，要么选，要么不选，不能只选一部分)。
目标是在不超过背包容量的前提下，选择一些物品放入背包，使得青背包中物品的总价值最大。
*/

// KnapsackII 函数用于解决 0 - 1 背包问题
// 参数 W 表示背包的容量
// 参数 weights 是一个切片，存储每个物品的重量
// 参数 values 是一个切片，存储每个物品的价值
// 返回值为在不超过背包容量的前提下能获得的最大价值
func KnapsackII(W int, weights []int, values []int) int {
	dp := make([]int, W+1)
	n := len(weights)
	for i := 0; i < n; i++ {
		for j := W; j >= weights[i]; j-- {
			dp[j] = max(dp[j], dp[j-weights[i]]+values[i])
		}
	}
	return dp[W]
}

/**
DP数组及下标的含义：
二维DP数组是根据上一行的已知信息计算当前行；
dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]] + values[i])
一维DP数组是将上一行的已知信息"拷贝"到当前行，直接在当前行进行计算。
dp[i][j] = max(dp[i][j], dp[i][j-weight[i]] + values[i])
如果是直接在当前行计算，就可以直接用一维数组来承载。
将i消去后，
- j：容量；
- dp[j]：容量为j的背包中物品的最大总价值；

递推公式：
dp[j] = max(dp[j], dp[j-weight[i]] + value[i])

初始化：
dp[0] = 0，表示容量为0时，无法装下任何物品。其余单元格的值都可以通过计算推导而来，无需显式初始化，且也不能初始化为其他值，可能会造成计算错误。
推导方向：
物品 -> 从前到后；
背包 -> 从后到前：
因为滚动数组的本质是屏蔽了对物品选取的表达，仅看当前容量为j时的背包情况。
当考虑物品i时，dp[0..W]的计算依赖于上一轮【考虑物品i-1时，计算&比较后的背包最大总价值】。上一轮的结果保存在当前行中，本轮同样基于当前行的信息进行重新计算。
如果是从前到后遍历背包，本轮计算会无情覆盖掉上一轮的计算结果，导致后面的单元格无法正确依赖上一轮的计算结果来进行推导。
比如，考虑物品i时，计算后更新了dp[j-1]，此单元格保存的结果由【上一轮对[0...i-1]物品任意选取且背包容量为j-1时的最大总价值】变为【本轮对[0...i]物品任意选取且背包容量为j-1时的最大总价值】。
此时若继续往后推导，i仍固定不变，j++，则dp[j]的计算依赖的j左边的单元格已被本轮的计算结果污染，在背包容量允许的情况下，会导致物品i被重复选取。
如果是从后往前遍历背包，则本轮的计算都仅依赖于上一轮计算的结果。
可举例推演一下做简单证明。
因为是从最大值开始逆向遍历背包的容量，且终止条件是当前物品的重量，所以不需要像二维DP数组一样判断当前物品是否能放入当前背包。
而且二维数组在放不下当前物品时，需要继承正上方的单元格的值，即dp[i][j] = dp[i-1][j]，表示当前容量下的最大总价值，还是等于考虑i-1个物品取/不取时的最大总价值。
而一维数组不需要任何显式继承赋值，因为当j < weight[i]时，也就是内层for循环没有处理的部分，本身已经是对i-1个物品取/不取时的最大总价值了。

遍历方向：
背包遍历方向：从后到前
物品&背包两个维度的for循环：是否可以颠倒？
若先物品再背包：
对第i个物品，算一遍各种容量下的背包最大总价值。再对第i+1个物品，基于取或不取i，算一遍各种容量下的背包最大总价值。以此类推……
每轮dp[0..W]的计算，是基于上一轮的dp[0..W]的。
满足了DP数组的定义：考虑物品i时，dp数组是物品i-1取/不取时，容量分别为[0...W]时的最大总价值。
若先背包再物品：
由上述推导方向可知，遍历方向应为 W -> 0，才能避免重复放入同一个物品。
对每个背包容量，计算dp[W...0]。
当容量为j时，考虑每个物品的放入情况。但由于dp[j]依赖的dp[0...j-1]此时还没基于每个物品的取用情况计算出来，无法依赖重叠子问题的计算结果来推导值。
会导致实际上每个容量的背包，都只能装一个物品。
也可举例推演一下做简单证明。
*/
