package dynamicprogramming

// 给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
//
// 示例 1：
//
// 输入：nums = [1,5,11,5]
// 输出：true
// 解释：数组可以分割成 [1, 5, 5] 和 [11] 。
//
// 示例 2：
//
// 输入：nums = [1,2,3,5]
// 输出：false
// 解释：数组不能分割成两个元素和相等的子集。
//
// 提示：
//
// 1 <= nums.length <= 200
// 1 <= nums[i] <= 100
func canPartition(nums []int) bool {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	dp := make([]int, target+1)
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return dp[target] == target
}

/**
分割成两个相等子集 -> 是否有一组元素的和 == 总和 / 2 -> 对每个元素考虑取/不取，且每个元素最多只能取一次 -> 0/1背包问题
元素值 == 物品重量 == 物品价值
本题转换为，容量为target的背包，是否正好能装下总价值为target的物品。

DP数组及下标含义：
- j：背包容量为j；
- dp[j]：背包容量为j时，背包所能装的物品的最大总价值。

递推公式：
dp[j] = max(dp[j], dp[j-nums[i]] + nums[i])
dp[j]总是<=j的，因为可能不取某个物品。当存在dp[j] == j，表示集合中的子集总和正好可以凑成总和j。

初始化：
dp[0] = 0

遍历方向：
先物品，再背包；
物品正序，背包倒序。

注意：
当sum为奇数时，无论如何也无法分割。如果此时不判断而直接除以2，向下取整后会有预期外的结果。
*/
