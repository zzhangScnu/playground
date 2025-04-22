package dynamicprogramming

// 有一堆石头，用整数数组 stones 表示。其中 stones[i] 表示第 i 块石头的重量。
//
// 每一回合，从中选出任意两块石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：
//
// 如果 x == y，那么两块石头都会被完全粉碎；
// 如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
//
// 最后，最多只会剩下一块 石头。返回此石头 最小的可能重量 。如果没有石头剩下，就返回 0。
//
// 示例 1：
//
// 输入：stones = [2,7,4,1,8,1]
// 输出：1
// 解释：
// 组合 2 和 4，得到 2，所以数组转化为 [2,7,1,8,1]，
// 组合 7 和 8，得到 1，所以数组转化为 [2,1,1,1]，
// 组合 2 和 1，得到 1，所以数组转化为 [1,1,1]，
// 组合 1 和 1，得到 0，所以数组转化为 [1]，这就是最优值。
//
// 示例 2：
//
// 输入：stones = [31,26,33,21,40]
// 输出：5
//
// 提示：
//
// 1 <= stones.length <= 30
// 1 <= stones[i] <= 100
func lastStoneWeightII(stones []int) int {
	var sum int
	for _, stone := range stones {
		sum += stone
	}
	target := sum / 2
	dp := make([]int, target+1)
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	return sum - dp[target]*2
}

/**
这题跟416-分割等和子集思路相似。

target := sum / 2，向下取整保证了动态规划求取的是较小或相等的那堆石头的重量总和。
dp[target]和target的关系是，target是目标凑成的重量和，dp[target]是实际凑成的重量和。
dp[j]表示容量为[0...target]时分别能装下的重量和。

dp[target]求出来以后，因为是较小的那堆，从而sum - dp[target]可得较大的那堆，较大-较小可得碰撞后的最小剩余重量。
*/
