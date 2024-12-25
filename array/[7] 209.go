package array

// 给定一个含有 n 个正整数的数组和一个正整数 target 。
//
// 找出该数组中满足其总和大于等于 target 的长度最小的 子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其
// 长度。如果不存在符合条件的子数组，返回 0 。
//
// 示例 1：
//
// 输入：target = 7, nums = [2,3,1,2,4,3]
// 输出：2
// 解释：子数组[4,3]是该条件下的长度最小的子数组。
//
// 示例 2：
//
// 输入：target = 4, nums = [1,4,4]
// 输出：1
//
// 示例 3：
//
// 输入：target = 11, nums = [1,1,1,1,1,1,1,1]
// 输出：0
//
// 提示：
//
// 1 <= target <= 10⁹
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁴
//
// 进阶：
//
// 如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n)) 时间复杂度的解法。
func minSubArrayLen(target int, nums []int) int {
	var i, val int
	res := len(nums) + 1
	for j := 0; j < len(nums); j++ {
		val += nums[j]
		for val >= target {
			if j-i+1 < res {
				res = j - i + 1
			}
			val -= nums[i]
			i++
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

/**
滑动窗口
是一种双指针的实现：
- 当不满足约束条件时，快指针推进，扩展窗口；
- 当满足约束条件时，慢指针推进，缩小窗口，同时更新结果。
注意更新结果的时机，需在满足约束条件时。否则无法兼容找不到结果的场景。

一些初始值：
慢指针&快指针：0
结果：数组长度+1，用来判断是否有被更新过、有无满足约束条件的场景
窗口区间：更新结果时，快慢指针都还未自增，所以是左闭右闭区间。
*/
