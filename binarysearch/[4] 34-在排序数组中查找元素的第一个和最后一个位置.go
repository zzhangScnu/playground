package binarysearch

// 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
//
// 如果数组中不存在目标值 target，返回 [-1, -1]。
//
// 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
//
// 示例 1：
//
// 输入：nums = [5,7,7,8,8,10], target = 8
// 输出：[3,4]
//
// 示例 2：
//
// 输入：nums = [5,7,7,8,8,10], target = 6
// 输出：[-1,-1]
//
// 示例 3：
//
// 输入：nums = [], target = 0
// 输出：[-1,-1]
//
// 提示：
//
// 0 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
// nums 是一个非递减数组
// -10⁹ <= target <= 10⁹
func searchRange(nums []int, target int) []int {
	return []int{
		searchLeftRange(nums, target), searchRightRange(nums, target),
	}
}

func searchLeftRange(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	if left > len(nums)-1 || nums[left] != target {
		return -1
	}
	return left
}

func searchRightRange(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}

/**
二分搜索的扩展：
- 前提：数组有序
- 找左边界(即比目标值小的元素数量)
	- 当找到目标值时，缩小右指针，向左逼近；
	- 循环结束时，需要检查左指针：
		- 如果索引已经超出数组长度，则代表：所有元素都比目标值小，左边界不存在；
		- 如果没超出且左指针指向的值不等于目标值，则代表：目标值在数组中不存在。
- 找右边界(即比目标值大的元素数量)
	- 当找到目标值时，缩小左指针，向右逼近。
	- 其他的特性，相当于左边界的镜像。
*/
