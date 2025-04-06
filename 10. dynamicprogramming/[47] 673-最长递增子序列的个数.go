package dynamicprogramming

// 给定一个未排序的整数数组
// nums ， 返回最长递增子序列的个数 。
//
// 注意 这个数列必须是 严格 递增的。
//
// 示例 1:
//
// 输入: [1,3,5,4,7]
// 输出: 2
// 解释: 有两个最长递增子序列，分别是 [1, 3, 4, 7] 和[1, 3, 5, 7]。
//
// 示例 2:
//
// 输入: [2,2,2,2,2]
// 输出: 5
// 解释: 最长递增子序列的长度是1，并且存在5个子序列的长度为1，因此输出5。
//
// 提示:
//
// 1 <= nums.length <= 2000
// -10⁶ <= nums[i] <= 10⁶
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] += dp[j]
			}
		}
	}
	return dp[n-1]
}

/**
DP数组及下标含义：
- i：当前游标指向位置；
- dp[i]：nums[0 ... i]间最长递增子序列的个数。

递推公式：
for i -> [0, n)
for j -> [0, i)
若 nums[j] < nums[i]，则表示nums[j ... i]单调递增，
	故dp[i] += dp[j]

初始化：
dp中每个单元格的值均为1，表示自身组成一个递增子序列。

遍历方向：
从左到右。
*/
