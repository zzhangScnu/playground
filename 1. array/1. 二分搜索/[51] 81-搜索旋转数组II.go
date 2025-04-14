package ___二分搜索

// 已知存在一个按非降序排列的整数数组 nums ，数组中的值不必互不相同。
//
// 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转 ，使数组变为 [nums[k],
// nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,
// 2,4,4,4,5,6,6,7] 在下标 5 处经旋转后可能变为 [4,5,6,6,7,0,1,2,4,4] 。
//
// 给你 旋转后 的数组 nums 和一个整数 target ，请你编写一个函数来判断给定的目标值是否存在于数组中。如果 nums 中存在这个目标值
// target ，则返回 true ，否则返回 false 。
//
// 你必须尽可能减少整个操作步骤。
//
// 示例 1：
//
// 输入：nums = [2,5,6,0,0,1,2], target = 0
// 输出：true
//
// 示例 2：
//
// 输入：nums = [2,5,6,0,0,1,2], target = 3
// 输出：false
//
// 提示：
//
// 1 <= nums.length <= 5000
// -10⁴ <= nums[i] <= 10⁴
// 题目数据保证 nums 在预先未知的某个下标上进行了旋转
// -10⁴ <= target <= 10⁴
//
// 进阶：
//
// 此题与 搜索旋转排序数组 相似，但本题中的 nums 可能包含 重复 元素。这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？
func searchY(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return true
		}
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left, right = left+1, right-1
			continue
		}
		if nums[left] <= nums[mid] {
			if nums[mid] < target || target < nums[left] {
				left = mid + 1
			} else if nums[mid] > target {
				right = mid - 1
			}
		} else {
			if nums[mid] > target || target > nums[right] {
				right = mid - 1
			} else if nums[mid] < target {
				left = mid + 1
			}
		}
	}
	return false
}

/**
跟33的区别在于，数组中存在重复元素。
这样会导致nums[mid]跟nums[0]及nums[n - 1]对比时，如果值相等，无法判断nums[mid]处于左侧数组还是右侧数组。
所以如果遇到这种场景且nums[mid]不是目标值，需要同时跳过nums[0]和nums[n - 1]，
即不断向内收缩left和right指针，跳过重复元素，创造可以识别左 / 右侧数组的条件。

注意：
if nums[left] <= nums[mid] { // 这里是nums[left]不是nums[0]，左侧边界元素可能被跳过了
	if nums[mid] < target || target < nums[left] { // 这里是nums[left]不是nums[0]，左侧边界元素可能被跳过了
		left = mid + 1
	} else if nums[mid] > target {
		right = mid - 1
	}
} else {
	if nums[mid] > target || target > nums[right] { 这里是nums[right]不是nums[len(nums)-1]，右侧边界元素可能被跳过了
		right = mid - 1
	} else if nums[mid] < target {
		left = mid + 1
	}
}
*/
