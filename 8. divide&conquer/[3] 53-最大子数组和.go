package divide_conquer

import "math"

// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
// 子数组 是数组中的一个连续部分。
//
// 示例 1：
//
// 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出：6
// 解释：连续子数组[4,-1,2,1] 的和最大，为6 。
//
// 示例 2：
//
// 输入：nums = [1]
// 输出：1
//
// 示例 3：
//
// 输入：nums = [5,4,-1,7,8]
// 输出：23
//
// 提示：
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
//
// 进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
func maxSubArray(nums []int) int {
	res := math.MinInt
	var maxCnt int
	for _, num := range nums {
		maxCnt += num
		if maxCnt > res {
			res = maxCnt
		}
		if maxCnt < 0 {
			maxCnt = 0
		}
	}
	return res
}

/**
if maxCnt > res {
	res = maxCnt
}
和
if maxCnt < 0 {
	maxCnt = 0
}
的先后顺序，影响全负数数组
*/
