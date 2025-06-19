package backtracking

import "slices"

// 给定一个整数数组 nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。
//
// 示例 1：
//
// 输入： nums = [4, 3, 2, 3, 5, 2, 1], k = 4
// 输出： True
// 说明： 有可能将其分成 4 个子集（5），（1,4），（2,3），（2,3）等于总和。
//
// 示例 2:
//
// 输入: nums = [1,2,3,4], k = 3
// 输出: false
//
// 提示：
//
// 1 <= k <= len(nums) <= 16
// 0 < nums[i] < 10000
// 每个元素的频率在 [1,4] 范围内
func canPartitionKSubsets(nums []int, k int) bool {
	if len(nums) < k {
		return false
	}
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	slices.Sort(nums)
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	bucketSum := make([]int, k)
	var traverse func(nums []int, numIndex int) bool
	traverse = func(nums []int, numIndex int) bool {
		if numIndex == len(nums) {
			return true
		}
		for i := 0; i < k; i++ {
			if bucketSum[i]+nums[numIndex] > target {
				continue
			}
			if i > 0 && bucketSum[i] == bucketSum[i-1] {
				continue
			}
			bucketSum[i] += nums[numIndex]
			if traverse(nums, numIndex+1) {
				return true
			}
			bucketSum[i] -= nums[numIndex]
			if bucketSum[i] == 0 {
				break // todo:???
			}
		}
		return false
	}
	return traverse(nums, 0)
}
