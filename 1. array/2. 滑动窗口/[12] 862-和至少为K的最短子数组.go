package ___滑动窗口

// 给你一个整数数组 nums 和一个整数 k ，找出 nums 中和至少为 k 的 最短非空子数组 ，并返回该子数组的长度。如果不存在这样的 子数组 ，返回
// -1 。
//
// 子数组 是数组中 连续 的一部分。
//
// 示例 1：
//
// 输入：nums = [1], k = 1
// 输出：1
//
// 示例 2：
//
// 输入：nums = [1,2], k = 4
// 输出：-1
//
// 示例 3：
//
// 输入：nums = [2,-1,2], k = 3
// 输出：3
//
// 提示：
//
// 1 <= nums.length <= 10⁵
// -10⁵ <= nums[i] <= 10⁵
// 1 <= k <= 10⁹

var MaxLen = 100_001

func shortestSubarray(nums []int, k int) int {
	minLen := MaxLen
	var windowRes int
	var left, right int
	for _, num := range nums {
		windowRes += num
		right++
		if windowRes >= k {
			windowRes -= k
			left++
			minLen = min(minLen, right-left+1)
		}
	}
	if minLen == MaxLen {
		return -1
	}
	return minLen
}
