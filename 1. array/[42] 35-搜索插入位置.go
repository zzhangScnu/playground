package array

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
// 请必须使用时间复杂度为 O(log n) 的算法。
//
// 示例 1:
//
// 输入: nums = [1,3,5,6], target = 5
// 输出: 2
//
// 示例 2:
//
// 输入: nums = [1,3,5,6], target = 2
// 输出: 1
//
// 示例 3:
//
// 输入: nums = [1,3,5,6], target = 7
// 输出: 4
//
// 提示:
//
// 1 <= nums.length <= 10⁴
// -10⁴ <= nums[i] <= 10⁴
// nums 为 无重复元素 的 升序 排列数组
// -10⁴ <= target <= 10⁴
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			if mid > 0 && nums[mid-1] < target {
				return mid
			}
			right = mid - 1
		}
	}
	if left == len(nums)-1 {
		return len(nums)
	}
	return -1
}

// first greater
