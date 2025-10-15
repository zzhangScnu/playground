package array

import "math"

// 给你一个整数数组 nums ，判断这个数组中是否存在长度为 3 的递增子序列。
//
// 如果存在这样的三元组下标 (i, j, k) 且满足 i < j < k ，使得 nums[i] < nums[j] < nums[k] ，返回
// true ；否则，返回 false 。
//
// 示例 1：
//
// 输入：nums = [1,2,3,4,5]
// 输出：true
// 解释：任何 i < j < k 的三元组都满足题意
//
// 示例 2：
//
// 输入：nums = [5,4,3,2,1]
// 输出：false
// 解释：不存在满足题意的三元组
//
// 示例 3：
//
// 输入：nums = [2,1,5,0,4,6]
// 输出：true
// 解释：其中一个满足题意的三元组是 (3, 4, 5)，因为 nums[3] == 0 < nums[4] == 4 < nums[5] == 6
//
// 提示：
//
// 1 <= nums.length <= 5 * 10⁵
// -2³¹ <= nums[i] <= 2³¹ - 1
//
// 进阶：你能实现时间复杂度为 O(n) ，空间复杂度为 O(1) 的解决方案吗？
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	first, second := nums[0], math.MaxInt64
	for _, num := range nums {
		if num > second {
			return true
		}
		if num > first {
			second = num
		} else if num < first {
			first = num
		}
	}
	return false
}

/**
思路：
维护2个变量，分别表示第一个数和第二个数。
寻找合适的数进行赋值，同时寻找第三个数。
当所有数都满足条件时，直接返回。
否则均为找到，返回false。

注意 first 和 second 的初始化值。
*/
