package dynamicprogramming

// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
//
// 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子
// 序列。
//
// 示例 1：
//
// 输入：nums = [10,9,2,5,3,7,101,18]
// 输出：4
// 解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
//
// 示例 2：
//
// 输入：nums = [0,1,0,3,2,3]
// 输出：4
//
// 示例 3：
//
// 输入：nums = [7,7,7,7,7,7,7]
// 输出：1
//
// 提示：
//
// 1 <= nums.length <= 2500
// -10⁴ <= nums[i] <= 10⁴
//
// 进阶：
//
// 你能将算法的时间复杂度降低到 O(n log(n)) 吗?
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	res := dp[0]
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

/**
DP数组及下标含义：
- i：当前游标指向索引为i；
- dp[i]：最长递增子序列的长度；

递推公式：
if nums[i] > nums[i-1],
dp[i] = max(dp[i], dp[i-1] + 1)

初始化：
可见i从i-1推导而来，
故dp[0] = 1

遍历方向：从左往右。
*/
