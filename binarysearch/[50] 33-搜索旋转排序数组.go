package binarysearch

// 整数数组 nums 按升序排列，数组中的值 互不相同 。
//
// 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[
// k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2
// ,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
//
// 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
//
// 你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
//
// 示例 1：
//
// 输入：nums = [4,5,6,7,0,1,2], target = 0
// 输出：4
//
// 示例 2：
//
// 输入：nums = [4,5,6,7,0,1,2], target = 3
// 输出：-1
//
// 示例 3：
//
// 输入：nums = [1], target = 0
// 输出：-1
//
// 提示：
//
// 1 <= nums.length <= 5000
// -10⁴ <= nums[i] <= 10⁴
// nums 中的每个值都 独一无二
// 题目数据保证 nums 在预先未知的某个下标上进行了旋转
// -10⁴ <= target <= 10⁴
func searchX(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if nums[mid] < target || target < nums[0] {
				left = mid + 1
			} else if nums[mid] > target {
				right = mid - 1
			}
		} else if nums[0] > nums[mid] {
			if nums[mid] > target || target > nums[len(nums)-1] {
				right = mid - 1
			} else if nums[mid] < target {
				left = mid + 1
			}
		}
	}
	return -1
}

/**
思路：

 ^              子数组A
 |             /
 |          /
 |       /
｜   /
｜/_ _ _ _ _ _ 子数组B
｜                  /
｜               /
————————————>
   0                    n-1

有序数组旋转后，右半部分桥接到左半部分前面，如上图所示，存在
若nums长度为n，
nums[0] > nums[n - 1]，即原本nums[n - 1]和nums[0]是相邻且递增的两个元素，经旋转后被分裂到A / B两个递增的子数组中。

所以本题需要分情况讨论：
- 若nums[mid]落在左侧子数组中，即nums[mid] >= nums[0]，
	- left = mid + 1，即需要往右搜索的场景：
		- 当target > nums[mid]，即目标值出现在坐标轴靠右；
		- 当需要搜索子数组B时，即target < nums[0]时；
 	- right = mid - 1，即需要往左搜索的场景：
		- 当target < nums[mid]，即目标值出现在坐标轴靠左；

- 若nums[mid]落在右侧子数组中，即nums[mid] < nums[0]，等价于nums[mid] <= nums[n - 1]，
	- right = mid - 1，即需要往左搜索的场景：
		- 当target < nums[mid]，即目标值出现在坐标轴靠左；
		- 当需要搜索子数组A时，当target > nums[n - 1]，等价于target >= nums[0]；
	- left = mid + 1，即需要往右搜索的场景：
		- 当target > nums[mid]，即目标值出现在坐标轴靠右。
*/
