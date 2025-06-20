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

/**
思路：数字视角
设有K个桶，每个桶的容量为target。
对于数组中的每一个数字，可以有K个选择，放入[0, K-1]区间的桶中。

if traverse(nums, numIndex+1) {
	return true
}
分支搜索，找到回溯树中任一符合条件路径即终止搜索，向上层层返回。

时间复杂度：
设 N = len(nums)，K = 桶的数量
每个数字有K个选择，则N个数字：
O(N^K)

跟桶视角相比可知，要尽量收缩可选择范围，因为每多一个选项都是指数级别增长；
相比之下只增加候选集个数如桶的数量，只是常数级别增长。
*/
