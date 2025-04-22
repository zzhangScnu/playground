package binarysearch

import "math"

// 峰值元素是指其值严格大于左右相邻值的元素。
//
// 给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
//
// 你可以假设 nums[-1] = nums[n] = -∞ 。
//
// 你必须实现时间复杂度为 O(log n) 的算法来解决此问题。
//
// 示例 1：
//
// 输入：nums = [1,2,3,1]
// 输出：2
// 解释：3 是峰值元素，你的函数应该返回其索引 2。
//
// 示例 2：
//
// 输入：nums = [1,2,1,3,5,6,4]
// 输出：1 或 5
// 解释：你的函数可以返回索引 1，其峰值元素为 2；
//
//	或者返回索引 5， 其峰值元素为 6。
//
// 提示：
//
// 1 <= nums.length <= 1000
// -2³¹ <= nums[i] <= 2³¹ - 1
// 对于所有有效的 i 都有 nums[i] != nums[i + 1]
func findPeakElement(nums []int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if get(nums, mid) > get(nums, mid-1) && get(nums, mid) > get(nums, mid+1) {
			return mid
		} else if get(nums, mid) < get(nums, mid+1) {
			low = mid + 1
		} else if get(nums, mid) < get(nums, mid-1) {
			high = mid - 1
		}
	}
	return -1
}

func get(nums []int, index int) int {
	if index < 0 || index >= len(nums) {
		return math.MinInt
	}
	return nums[index]
}

/**
思路：
在一个波浪图形中，找出任意一个峰值。

技巧：
用get()函数屏蔽掉边界处理，避免额外的特例判断。

因为要寻找峰值，且有约束对于所有有效的i都有 nums[i] != nums[i + 1]，意味着单调递增 => 峰值 => 单调递减，
故峰值左右都是有序的，且题目要求O(logN)的时间复杂度，自然引入二分搜索。
每次通过对nums[mid]处于上坡 / 下坡的判断，排除掉一半搜索区间，比线性查找更高效。
*/
